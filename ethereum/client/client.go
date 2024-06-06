package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dan13ram/wpokt-oracle/models"

	log "github.com/sirupsen/logrus"
)

const (
	MaxQueryBlocks uint64 = 100000
)

type EthereumClient interface {
	ValidateNetwork() error
	GetBlockHeight() (uint64, error)
	GetChainID() (*big.Int, error)
	GetClient() *ethclient.Client
	GetTransactionByHash(txHash string) (*types.Transaction, bool, error)
	GetTransactionReceipt(txHash string) (*types.Receipt, error)
}

type ethereumClient struct {
	Timeout   time.Duration
	ChainID   uint64
	ChainName string

	client *ethclient.Client

	logger *log.Entry
}

func (c *ethereumClient) GetClient() *ethclient.Client {
	return c.client
}

func (c *ethereumClient) GetBlockHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	blockNumber, err := c.client.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return blockNumber, nil
}

func (c *ethereumClient) GetChainID() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	chainId, err := c.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	return chainId, nil
}

func (c *ethereumClient) ValidateNetwork() error {
	c.logger.Debugf("Validating network")

	chainID, err := c.GetChainID()
	if err != nil {
		return fmt.Errorf("failed to validate network: %s", err)
	}
	if chainID.Cmp(big.NewInt(int64(c.ChainID))) != 0 {
		return fmt.Errorf("failed to validate network: expected chain id %d, got %s", c.ChainID, chainID)
	}

	c.logger.Debugf("Validated network")
	return nil
}

func (c *ethereumClient) GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	tx, isPending, err := c.client.TransactionByHash(ctx, common.HexToHash(txHash))
	return tx, isPending, err
}

func (c *ethereumClient) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	receipt, err := c.client.TransactionReceipt(ctx, common.HexToHash(txHash))
	return receipt, err
}

func NewClient(config models.EthereumNetworkConfig) (EthereumClient, error) {
	logger := log.
		WithField("module", "ethereum").
		WithField("package", "client").
		WithField("chain_name", strings.ToLower(config.ChainName)).
		WithField("chain_id", config.ChainID)
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		logger.WithError(err).Error("failed to connect to rpc")
		return nil, fmt.Errorf("failed to connect to rpc")
	}

	ethclient := &ethereumClient{
		Timeout:   time.Duration(config.TimeoutMS) * time.Millisecond,
		ChainID:   config.ChainID,
		ChainName: config.ChainName,

		client: client,

		logger: logger,
	}

	err = ethclient.ValidateNetwork()
	if err != nil {
		logger.WithError(err).Error("failed to validate network")
		return nil, fmt.Errorf("failed to validate network")
	}

	return ethclient, err
}
