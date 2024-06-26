package db

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/dan13ram/wpokt-oracle/common"
	mocks "github.com/dan13ram/wpokt-oracle/db/mocks"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/dan13ram/wpokt-oracle/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageTestSuite struct {
	suite.Suite
	mockDB     *mocks.MockDatabase
	oldMongoDB Database
	db         MessageDB
}

func (suite *MessageTestSuite) SetupTest() {
	suite.mockDB = mocks.NewMockDatabase(suite.T())
	suite.oldMongoDB = mongoDB
	mongoDB = suite.mockDB
	suite.db = &messageDB{}
}

func (suite *MessageTestSuite) TearDownTest() {
	mongoDB = suite.oldMongoDB
}

func (suite *MessageTestSuite) TestNewMessageBody() {
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	amount := big.NewInt(100)

	messageBody, err := suite.db.NewMessageBody(senderAddress[:], amount, recipientAddress[:])
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), strings.ToLower(senderAddress.Hex()), messageBody.SenderAddress)
	assert.Equal(suite.T(), amount.String(), messageBody.Amount)
	assert.Equal(suite.T(), strings.ToLower(recipientAddress.Hex()), messageBody.RecipientAddress)
}

func (suite *MessageTestSuite) TestNewMessageBody_InvalidSender() {
	senderAddress := []byte{1, 2, 3}
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	amount := big.NewInt(100)
	expectedError := fmt.Errorf("invalid sender address: %w", common.ErrInvalidAddressLength)

	_, err := suite.db.NewMessageBody(senderAddress[:], amount, recipientAddress[:])
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *MessageTestSuite) TestNewMessageBody_InvalidRecipient() {
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	recipientAddress := []byte{4, 5, 6}
	amount := big.NewInt(100)
	expectedError := fmt.Errorf("invalid recipient address: %w", common.ErrInvalidAddressLength)

	_, err := suite.db.NewMessageBody(senderAddress[:], amount, recipientAddress[:])
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *MessageTestSuite) TestNewMessageContent() {
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	destinationDomain := uint32(2)
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	messageBody := models.MessageBody{
		SenderAddress:    "0x010203",
		Amount:           big.NewInt(100).String(),
		RecipientAddress: "0x040506",
	}

	messageContent, err := suite.db.NewMessageContent(nonce, originDomain, senderAddress[:], destinationDomain, recipientAddress[:], messageBody)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), common.HyperlaneVersion, messageContent.Version)
	assert.Equal(suite.T(), nonce, messageContent.Nonce)
	assert.Equal(suite.T(), originDomain, messageContent.OriginDomain)
	assert.Equal(suite.T(), strings.ToLower(senderAddress.Hex()), messageContent.Sender)
	assert.Equal(suite.T(), destinationDomain, messageContent.DestinationDomain)
	assert.Equal(suite.T(), strings.ToLower(recipientAddress.Hex()), messageContent.Recipient)
	assert.Equal(suite.T(), messageBody, messageContent.MessageBody)
}

func (suite *MessageTestSuite) TestNewMessageContent_InvalidSender() {
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := []byte{1, 2, 3}
	destinationDomain := uint32(2)
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	messageBody := models.MessageBody{
		SenderAddress:    "0x010203",
		Amount:           big.NewInt(100).String(),
		RecipientAddress: "0x040506",
	}
	expectedError := fmt.Errorf("invalid sender address: %w", common.ErrInvalidAddressLength)

	_, err := suite.db.NewMessageContent(nonce, originDomain, senderAddress, destinationDomain, recipientAddress[:], messageBody)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *MessageTestSuite) TestNewMessageContent_InvalidRecipient() {
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	destinationDomain := uint32(2)
	recipientAddress := []byte{4, 5, 6}
	messageBody := models.MessageBody{
		SenderAddress:    "0x010203",
		Amount:           big.NewInt(100).String(),
		RecipientAddress: "0x040506",
	}
	expectedError := fmt.Errorf("invalid recipient address: %w", common.ErrInvalidAddressLength)

	_, err := suite.db.NewMessageContent(nonce, originDomain, senderAddress[:], destinationDomain, recipientAddress, messageBody)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *MessageTestSuite) TestNewMessage() {
	txDoc := &models.Transaction{
		ID:   &primitive.ObjectID{},
		Hash: "0x123",
	}
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	destinationDomain := uint32(2)
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	messageBody := models.MessageBody{
		SenderAddress:    strings.ToLower(senderAddress.Hex()),
		Amount:           big.NewInt(100).String(),
		RecipientAddress: strings.ToLower(recipientAddress.Hex()),
	}

	content, _ := suite.db.NewMessageContent(nonce, originDomain, senderAddress[:], destinationDomain, recipientAddress[:], messageBody)
	status := models.MessageStatusPending

	message, err := suite.db.NewMessage(txDoc, content, status)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), *txDoc.ID, message.OriginTransaction)
	assert.Equal(suite.T(), txDoc.Hash, message.OriginTransactionHash)
	assert.Equal(suite.T(), status, message.Status)
}

func (suite *MessageTestSuite) TestNewMessage_InvalidTxDoc() {
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	destinationDomain := uint32(2)
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	messageBody := models.MessageBody{
		SenderAddress:    strings.ToLower(senderAddress.Hex()),
		Amount:           big.NewInt(100).String(),
		RecipientAddress: strings.ToLower(recipientAddress.Hex()),
	}

	content, _ := suite.db.NewMessageContent(nonce, originDomain, senderAddress[:], destinationDomain, recipientAddress[:], messageBody)
	status := models.MessageStatusPending

	expectedError := fmt.Errorf("invalid txDoc")

	_, err := suite.db.NewMessage(nil, content, status)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *MessageTestSuite) TestNewMessage_InvalidMessageID() {
	nonce := uint32(1)
	originDomain := uint32(1)
	senderAddress := ethcommon.BytesToAddress([]byte{1, 2, 3})
	destinationDomain := uint32(2)
	recipientAddress := ethcommon.BytesToAddress([]byte{4, 5, 6})
	messageBody := models.MessageBody{
		SenderAddress:    "0x12334",
		Amount:           big.NewInt(100).String(),
		RecipientAddress: strings.ToLower(recipientAddress.Hex()),
	}

	content, _ := suite.db.NewMessageContent(nonce, originDomain, senderAddress[:], destinationDomain, recipientAddress[:], messageBody)
	status := models.MessageStatusPending

	txDoc := &models.Transaction{
		ID:   &primitive.ObjectID{},
		Hash: "0x123",
	}
	_, err := suite.db.NewMessage(txDoc, content, status)
	assert.Error(suite.T(), err)
}

func (suite *MessageTestSuite) TestFindMessage() {
	filter := bson.M{"_id": primitive.NewObjectID()}
	expectedMessage := models.Message{}

	suite.mockDB.EXPECT().FindOne(common.CollectionMessages, filter, &expectedMessage).Return(nil).Once()

	gotMessage, err := suite.db.FindMessage(filter)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedMessage, gotMessage)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestUpdateMessage() {
	messageID := primitive.NewObjectID()
	update := bson.M{"status": models.MessageStatusSigned}

	suite.mockDB.EXPECT().UpdateOne(common.CollectionMessages, bson.M{"_id": &messageID}, bson.M{"$set": update}).Return(primitive.ObjectID{}, nil).Once()

	err := suite.db.UpdateMessage(&messageID, update)
	assert.NoError(suite.T(), err)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestUpdateMessage_NilMessageID() {
	update := bson.M{"status": models.MessageStatusSigned}

	err := suite.db.UpdateMessage(nil, update)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), fmt.Errorf("messageID is nil"), err)
}

func (suite *MessageTestSuite) TestUpdateMessageByMessageID() {
	messageID := [32]byte{}
	update := bson.M{"status": models.MessageStatusSigned}
	messageIDHex := common.Ensure0xPrefix(common.HexFromBytes(messageID[:]))

	suite.mockDB.EXPECT().UpdateOne(common.CollectionMessages, bson.M{"message_id": messageIDHex}, bson.M{"$set": update}).Return(primitive.ObjectID{}, nil).Once()

	_, err := suite.db.UpdateMessageByMessageID(messageID, update)
	assert.NoError(suite.T(), err)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestInsertMessage() {
	message := models.Message{
		ID: &primitive.ObjectID{},
	}
	insertedID := primitive.NewObjectID()

	suite.mockDB.EXPECT().InsertOne(common.CollectionMessages, message).Return(insertedID, nil).Once()

	gotID, err := suite.db.InsertMessage(message)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestInsertMessage_DuplicateKeyError() {
	message := models.Message{
		OriginTransactionHash: "0x123",
	}
	duplicateError := mongo.WriteError{Code: 11000}
	insertedID := primitive.NewObjectID()
	existingMessage := models.Message{
		ID: &insertedID,
	}

	suite.mockDB.EXPECT().InsertOne(common.CollectionMessages, message).Return(primitive.ObjectID{}, duplicateError).Once()
	suite.mockDB.EXPECT().FindOne(common.CollectionMessages, bson.M{"origin_transaction_hash": message.OriginTransactionHash}, &models.Message{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*models.Message)
		*arg = existingMessage
	})

	gotID, err := suite.db.InsertMessage(message)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestInsertMessage_DuplicateKeyError_FindError() {
	message := models.Message{
		OriginTransactionHash: "0x123",
	}
	duplicateError := mongo.WriteError{Code: 11000}
	insertedID := primitive.NewObjectID()
	expectedError := fmt.Errorf("find error")

	suite.mockDB.EXPECT().InsertOne(common.CollectionMessages, message).Return(insertedID, duplicateError).Once()
	suite.mockDB.EXPECT().FindOne(common.CollectionMessages, bson.M{"origin_transaction_hash": message.OriginTransactionHash}, &models.Message{}).Return(expectedError).Once()

	gotID, err := suite.db.InsertMessage(message)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestInsertMessage_InsertError() {
	message := models.Message{
		OriginTransactionHash: "0x123",
	}
	insertedID := primitive.NewObjectID()
	expectedError := fmt.Errorf("insert error")

	suite.mockDB.EXPECT().InsertOne(common.CollectionMessages, message).Return(insertedID, expectedError).Once()

	gotID, err := suite.db.InsertMessage(message)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestGetPendingMessages() {
	signerToExclude := "signer1"
	chain := models.Chain{ChainDomain: 1}
	messages := []models.Message{
		{
			ID:     &primitive.ObjectID{},
			Status: models.MessageStatusPending,
		},
	}
	filter := bson.M{
		"$and": []bson.M{
			{"content.destination_domain": chain.ChainDomain},
			{"$or": []bson.M{
				{"status": models.MessageStatusPending},
				{"status": models.MessageStatusSigned},
			}},
			{"$nor": []bson.M{
				{"signatures": bson.M{
					"$elemMatch": bson.M{"signer": signerToExclude},
				}},
			}},
		},
	}
	sort := bson.M{"content.nonce": 1}

	suite.mockDB.EXPECT().FindManySorted(common.CollectionMessages, filter, sort, &[]models.Message{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(3).(*[]models.Message)
		*arg = messages
	})

	gotMessages, err := suite.db.GetPendingMessages(signerToExclude, chain)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), messages, gotMessages)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestGetSignedMessages() {
	chain := models.Chain{ChainDomain: 1}
	messages := []models.Message{
		{
			ID:     &primitive.ObjectID{},
			Status: models.MessageStatusSigned,
		},
	}
	sort := bson.M{"sequence": 1}
	filter := bson.M{
		"content.destination_domain": chain.ChainDomain,
		"status":                     models.MessageStatusSigned,
	}

	suite.mockDB.EXPECT().FindManySorted(common.CollectionMessages, filter, sort, &[]models.Message{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(3).(*[]models.Message)
		*arg = messages
	})

	gotMessages, err := suite.db.GetSignedMessages(chain)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), messages, gotMessages)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *MessageTestSuite) TestGetBroadcastedMessages() {
	chain := models.Chain{ChainDomain: 1}
	messages := []models.Message{
		{
			ID:     &primitive.ObjectID{},
			Status: models.MessageStatusBroadcasted,
		},
	}
	filter := bson.M{
		"content.destination_domain": chain.ChainDomain,
		"status":                     models.MessageStatusBroadcasted,
		"transaction":                nil,
	}

	suite.mockDB.EXPECT().FindMany(common.CollectionMessages, filter, &[]models.Message{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*[]models.Message)
		*arg = messages
	})

	gotMessages, err := suite.db.GetBroadcastedMessages(chain)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), messages, gotMessages)
	suite.mockDB.AssertExpectations(suite.T())
}

func TestMessageTestSuite(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}
