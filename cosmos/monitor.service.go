package cosmos

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	crypto "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"go.mongodb.org/mongo-driver/bson"

	log "github.com/sirupsen/logrus"

	"github.com/dan13ram/wpokt-oracle/common"
	cosmos "github.com/dan13ram/wpokt-oracle/cosmos/client"
	"github.com/dan13ram/wpokt-oracle/db"
	"github.com/dan13ram/wpokt-oracle/models"
	"github.com/dan13ram/wpokt-oracle/service"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CosmosMessageMonitorRunnable struct {
	multisigAddressBytes []byte

	mintControllerMap         map[uint32][]byte
	supportedChainIDsEthereum map[uint32]bool

	chain  models.Chain
	config models.CosmosNetworkConfig
	client cosmos.CosmosClient

	logger *log.Entry

	startBlockHeight   uint64
	currentBlockHeight uint64

	db db.DB
}

func (x *CosmosMessageMonitorRunnable) Run() {
	x.UpdateCurrentHeight()
	x.SyncNewTxs()
	x.ConfirmTxs()
	x.CreateRefundsOrMessagesForConfirmedTxs()
}

func (x *CosmosMessageMonitorRunnable) Height() uint64 {
	return uint64(x.currentBlockHeight)
}

func (x *CosmosMessageMonitorRunnable) UpdateCurrentHeight() {
	height, err := x.client.GetLatestBlockHeight()
	if err != nil {
		x.logger.
			WithError(err).
			Error("could not get current block height")
		return
	}
	x.currentBlockHeight = uint64(height)
	x.logger.
		WithField("current_block_height", x.currentBlockHeight).
		Info("updated current block height")
}

func (x *CosmosMessageMonitorRunnable) CreateRefund(
	txRes *sdk.TxResponse,
	txDoc *models.Transaction,
	toAddr []byte,
	amount sdk.Coin,
) bool {

	refund, err := x.db.NewRefund(txRes, txDoc, toAddr, amount)
	if err != nil {
		x.logger.WithError(err).Errorf("Error creating refund")
		return false
	}

	insertedID, err := x.db.InsertRefund(refund)
	if err != nil {
		x.logger.WithError(err).Errorf("Error inserting refund")
		return false
	}

	err = x.db.UpdateTransaction(txDoc.ID, bson.M{"refund": insertedID})
	if err != nil {
		x.logger.WithError(err).Errorf("Error updating transaction")
		return false
	}

	return true
}

func (x *CosmosMessageMonitorRunnable) CreateMessage(
	txRes *sdk.TxResponse,
	tx *tx.Tx,
	txDoc *models.Transaction,
	senderAddress []byte,
	amountCoin sdk.Coin,
	memo models.MintMemo,
) bool {
	recipientAddr, err := common.BytesFromAddressHex(memo.Address)
	if err != nil {
		x.logger.WithError(err).Errorf("Error parsing recipient address")
		return false
	}

	messageBody, err := x.db.NewMessageBody(
		senderAddress,
		amountCoin.Amount.BigInt(),
		recipientAddr,
	)
	if err != nil {
		x.logger.WithError(err).Errorf("Error creating message body")
		return false
	}

	if len(tx.AuthInfo.SignerInfos) == 0 {
		x.logger.Errorf("No signer infos found")
		return false
	}
	nonce := uint32(tx.AuthInfo.SignerInfos[0].Sequence)
	originDomain := uint32(x.chain.ChainDomain)
	chainID, _ := strconv.Atoi(memo.ChainID)
	destinationDomain := uint32(chainID)
	destMintController, ok := x.mintControllerMap[destinationDomain]
	if !ok {
		x.logger.Errorf("Mint controller not found")
		return false
	}

	messageContent, err := x.db.NewMessageContent(
		nonce,
		originDomain,
		senderAddress,
		destinationDomain,
		destMintController,
		messageBody,
	)
	if err != nil {
		x.logger.WithError(err).Errorf("Error creating message content")
		return false
	}

	message, err := x.db.NewMessage(txDoc, messageContent, models.MessageStatusPending)
	if err != nil {
		x.logger.WithError(err).Errorf("Error creating message")
		return false
	}

	messageID, err := x.db.InsertMessage(message)
	if err != nil {
		x.logger.WithError(err).Errorf("Error inserting message")
		return false
	}

	txDoc.Messages = append(txDoc.Messages, messageID)

	err = x.db.UpdateTransaction(txDoc.ID, bson.M{"messages": common.RemoveDuplicates(txDoc.Messages)})
	if err != nil {
		x.logger.WithError(err).Errorf("Error updating transaction")
		return false
	}

	return true
}

func (x *CosmosMessageMonitorRunnable) SyncNewTxs() bool {
	x.logger.Infof("Syncing new txs")
	if x.currentBlockHeight <= x.startBlockHeight {
		x.logger.Infof("No new blocks to sync")
		return true
	}

	txResponses, err := x.client.GetTxsSentToAddressAfterHeight(x.config.MultisigAddress, x.startBlockHeight)
	if err != nil {
		x.logger.WithError(err).Errorf("Error getting new txs")
		return false
	}
	x.logger.Infof("Found %d txs to sync", len(txResponses))
	success := true
	for _, txResponse := range txResponses {
		logger := x.logger.WithField("tx_hash", txResponse.TxHash).WithField("section", "sync")

		result, err := utilValidateTxToCosmosMultisig(txResponse, x.config, x.supportedChainIDsEthereum, x.currentBlockHeight)
		if err != nil {
			logger.WithError(err).Errorf("Error validating tx")
			success = false
			continue
		}

		transaction, err := x.db.NewCosmosTransaction(txResponse, x.chain, result.SenderAddress, x.multisigAddressBytes, result.TxStatus)
		if err != nil {
			logger.WithError(err).Errorf("Error creating transaction")
			success = false
			continue
		}

		_, err = x.db.InsertTransaction(transaction)
		if err != nil {
			logger.WithError(err).Errorf("Error inserting transaction")
			success = false
			continue
		}
	}

	if success {
		x.startBlockHeight = x.currentBlockHeight
	}

	return success
}

func (x *CosmosMessageMonitorRunnable) ValidateAndConfirmTx(txDoc *models.Transaction) bool {
	logger := x.logger.WithField("tx_hash", txDoc.Hash).WithField("section", "confirm")
	txResponse, err := x.client.GetTx(txDoc.Hash)
	if err != nil {
		logger.WithError(err).Errorf("Error getting tx")
		return false
	}

	result, err := utilValidateTxToCosmosMultisig(txResponse, x.config, x.supportedChainIDsEthereum, x.currentBlockHeight)
	if err != nil {
		logger.WithError(err).Errorf("Error validating tx")
		return false
	}

	update := bson.M{
		"confirmations": result.Confirmations,
		"status":        result.TxStatus,
	}
	err = x.db.UpdateTransaction(txDoc.ID, update)
	if err != nil {
		logger.WithError(err).Errorf("Error updating transaction")
		return false
	}
	return true
}

func (x *CosmosMessageMonitorRunnable) ConfirmTxs() bool {
	x.logger.Infof("Confirming txs")
	txs, err := x.db.GetPendingTransactionsTo(x.chain, x.multisigAddressBytes)
	if err != nil {
		x.logger.WithError(err).Errorf("Error getting pending txs")
		return false
	}
	x.logger.Infof("Found %d pending txs", len(txs))
	success := true
	for _, txDoc := range txs {
		success = x.ValidateAndConfirmTx(&txDoc) && success
	}

	return success
}

func (x *CosmosMessageMonitorRunnable) ValidateTxAndCreate(txDoc *models.Transaction) bool {
	logger := x.logger.WithField("tx_hash", txDoc.Hash).WithField("section", "create")
	txResponse, err := x.client.GetTx(txDoc.Hash)
	if err != nil {
		logger.WithError(err).Errorf("Error getting tx")
		return false
	}

	result, err := utilValidateTxToCosmosMultisig(txResponse, x.config, x.supportedChainIDsEthereum, x.currentBlockHeight)
	if err != nil {
		logger.WithError(err).Errorf("Error validating tx")
		return false
	}

	if result.TxStatus == models.TransactionStatusPending {
		logger.Debugf("Found tx with status pending")
		return false
	}

	if result.TxStatus != models.TransactionStatusConfirmed {
		logger.Warnf("Found tx with status %s", result.TxStatus)
		err = x.db.UpdateTransaction(txDoc.ID, bson.M{"status": result.TxStatus})
		if err != nil {
			logger.WithError(err).Errorf("Error updating transaction")
			return false
		}
		return true
	}

	if lockID, err := x.db.LockWriteTransaction(txDoc); err != nil {
		logger.WithError(err).Errorf("Error locking transaction")
		return false
	} else {
		//nolint:errcheck
		defer x.db.Unlock(lockID)
	}

	if result.NeedsRefund {
		return x.CreateRefund(txResponse, txDoc, result.SenderAddress, result.Amount)
	}

	return x.CreateMessage(txResponse, result.Tx, txDoc, result.SenderAddress, result.Amount, result.Memo)
}

func (x *CosmosMessageMonitorRunnable) CreateRefundsOrMessagesForConfirmedTxs() bool {
	x.logger.Infof("Creating refunds or messages for confirmed txs")
	txDocs, err := x.db.GetConfirmedTransactionsTo(x.chain, x.multisigAddressBytes)
	if err != nil {
		x.logger.WithError(err).Errorf("Error getting confirmed txs")
		return false
	}
	x.logger.Infof("Found %d confirmed txs", len(txDocs))
	success := true
	for _, txDoc := range txDocs {
		success = x.ValidateTxAndCreate(&txDoc) && success
	}

	return success
}

func (x *CosmosMessageMonitorRunnable) InitStartBlockHeight(lastHealth *models.RunnerServiceStatus) {
	if lastHealth == nil || lastHealth.BlockHeight == 0 {
		x.logger.Debugf("Invalid last health")
	} else {
		x.logger.Debugf("Last block height: %d", lastHealth.BlockHeight)
		x.startBlockHeight = lastHealth.BlockHeight
	}
	if x.startBlockHeight == 0 {
		x.logger.Debugf("Start block height is zero")
		x.startBlockHeight = x.currentBlockHeight
	} else if x.startBlockHeight > x.currentBlockHeight {
		x.logger.Debugf("Start block height is greater than current block height")
		x.startBlockHeight = x.currentBlockHeight
	}
	x.logger.Infof("Initialized start block height: %d", x.startBlockHeight)
}

var dbNewDB = db.NewDB
var cosmosNewClient = cosmos.NewClient

func NewMessageMonitor(
	config models.CosmosNetworkConfig,
	mintControllerMap map[uint32][]byte,
	ethNetworks []models.EthereumNetworkConfig,
	lastHealth *models.RunnerServiceStatus,
) service.Runnable {
	logger := log.
		WithField("module", "cosmos").
		WithField("service", "monitor").
		WithField("chain_name", strings.ToLower(config.ChainName)).
		WithField("chain_id", strings.ToLower(config.ChainID))

	if !config.MessageMonitor.Enabled {
		logger.Fatalf("Message monitor is not enabled")
	}

	logger.Debugf("Initializing")

	var pks []crypto.PubKey
	for _, pk := range config.MultisigPublicKeys {
		pKey, err := common.CosmosPublicKeyFromHex(pk)
		if err != nil {
			logger.Fatalf("Error parsing public key: %s", err)
		}
		pks = append(pks, pKey)
	}

	multisigPk := multisig.NewLegacyAminoPubKey(int(config.MultisigThreshold), pks)
	multisigAddress, _ := common.Bech32FromBytes(config.Bech32Prefix, multisigPk.Address().Bytes())

	if !strings.EqualFold(multisigAddress, config.MultisigAddress) {
		logger.Fatalf("Multisig address does not match config")
	}

	client, err := cosmosNewClient(config)
	if err != nil {
		logger.Fatalf("Error creating cosmos client: %s", err)
	}

	supportedChainIDsEthereum := make(map[uint32]bool)
	for _, ethNetwork := range ethNetworks {
		supportedChainIDsEthereum[uint32(ethNetwork.ChainID)] = true
	}

	// TODO: check max amount for corresponding chain and disallow if too high

	x := &CosmosMessageMonitorRunnable{
		multisigAddressBytes: multisigPk.Address().Bytes(),

		startBlockHeight:   config.StartBlockHeight,
		currentBlockHeight: 0,
		client:             client,

		mintControllerMap:         mintControllerMap,
		supportedChainIDsEthereum: supportedChainIDsEthereum,

		chain:  utilParseChain(config),
		config: config,

		logger: logger,

		db: dbNewDB(),
	}

	x.UpdateCurrentHeight()

	x.InitStartBlockHeight(lastHealth)

	logger.Infof("Initialized")

	return x
}
