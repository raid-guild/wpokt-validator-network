package db

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dan13ram/wpokt-oracle/common"
	mocks "github.com/dan13ram/wpokt-oracle/db/mocks"
	"github.com/dan13ram/wpokt-oracle/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type RefundTestSuite struct {
	suite.Suite
	oldMongoDB Database
	mockDB     *mocks.MockDatabase
}

func (suite *RefundTestSuite) SetupTest() {
	suite.mockDB = mocks.NewMockDatabase(suite.T())
	suite.oldMongoDB = mongoDB
	mongoDB = suite.mockDB
}

func (suite *RefundTestSuite) TearDownTest() {
	mongoDB = suite.oldMongoDB
}

func (suite *RefundTestSuite) TestNewRefund() {
	txRes := &sdk.TxResponse{TxHash: "0x010203"}
	txDoc := &models.Transaction{
		ID:   &primitive.ObjectID{},
		Hash: "0x010203",
	}
	recipientAddress := ethcommon.HexToAddress("0x010203")
	amountCoin := sdk.Coin{Amount: math.NewInt(100)}

	refund, err := NewRefund(txRes, txDoc, recipientAddress[:], amountCoin)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), txDoc.ID, &refund.OriginTransaction)
	assert.Equal(suite.T(), txDoc.Hash, refund.OriginTransactionHash)
	assert.Equal(suite.T(), recipientAddress.Hex(), refund.Recipient)
	assert.Equal(suite.T(), uint64(100), refund.Amount)
	assert.Equal(suite.T(), models.RefundStatusPending, refund.Status)
}

func (suite *RefundTestSuite) TestInsertRefund() {
	refund := models.Refund{
		ID: &primitive.ObjectID{},
	}
	insertedID := primitive.NewObjectID()

	suite.mockDB.On("InsertOne", common.CollectionRefunds, refund).Return(insertedID, nil).Once()

	gotID, err := InsertRefund(refund)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RefundTestSuite) TestInsertRefund_DuplicateKeyError() {
	refund := models.Refund{
		OriginTransactionHash: "0x123",
	}
	duplicateError := mongo.WriteError{Code: 11000}
	insertedID := primitive.NewObjectID()
	existingRefund := models.Refund{
		ID: &insertedID,
	}

	suite.mockDB.On("InsertOne", common.CollectionRefunds, refund).Return(primitive.ObjectID{}, duplicateError).Once()
	suite.mockDB.On("FindOne", common.CollectionRefunds, bson.M{"origin_transaction_hash": refund.OriginTransactionHash}, &models.Refund{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*models.Refund)
		*arg = existingRefund
	})

	gotID, err := InsertRefund(refund)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), insertedID, gotID)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RefundTestSuite) TestUpdateRefund() {
	refundID := primitive.NewObjectID()
	update := bson.M{"status": models.RefundStatusSigned}

	suite.mockDB.On("UpdateOne", common.CollectionRefunds, bson.M{"_id": &refundID}, bson.M{"$set": update}).Return(primitive.ObjectID{}, nil).Once()

	err := UpdateRefund(&refundID, update)
	assert.NoError(suite.T(), err)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RefundTestSuite) TestGetPendingRefunds() {
	signerToExclude := "signer1"
	refunds := []models.Refund{
		{
			ID:     &primitive.ObjectID{},
			Status: models.RefundStatusPending,
		},
	}
	filter := bson.M{
		"$and": []bson.M{
			{"$or": []bson.M{
				{"status": models.RefundStatusPending},
				{"status": models.RefundStatusSigned},
			}},
			{"$nor": []bson.M{
				{"signatures": bson.M{
					"$elemMatch": bson.M{"signer": signerToExclude},
				}},
			}},
		},
	}

	suite.mockDB.On("FindMany", common.CollectionRefunds, filter, &[]models.Refund{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*[]models.Refund)
		*arg = refunds
	})

	gotRefunds, err := GetPendingRefunds(signerToExclude)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), refunds, gotRefunds)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RefundTestSuite) TestGetSignedRefunds() {
	refunds := []models.Refund{
		{
			ID:     &primitive.ObjectID{},
			Status: models.RefundStatusSigned,
		},
	}
	filter := bson.M{"status": models.RefundStatusSigned}
	sort := bson.M{"sequence": 1}

	suite.mockDB.On("FindManySorted", common.CollectionRefunds, filter, sort, &[]models.Refund{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(3).(*[]models.Refund)
		*arg = refunds
	})

	gotRefunds, err := GetSignedRefunds()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), refunds, gotRefunds)
	suite.mockDB.AssertExpectations(suite.T())
}

func (suite *RefundTestSuite) TestGetBroadcastedRefunds() {
	refunds := []models.Refund{
		{
			ID:     &primitive.ObjectID{},
			Status: models.RefundStatusBroadcasted,
		},
	}
	filter := bson.M{"status": models.RefundStatusBroadcasted, "transaction": nil}

	suite.mockDB.On("FindMany", common.CollectionRefunds, filter, &[]models.Refund{}).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*[]models.Refund)
		*arg = refunds
	})

	gotRefunds, err := GetBroadcastedRefunds()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), refunds, gotRefunds)
	suite.mockDB.AssertExpectations(suite.T())
}

func TestRefundTestSuite(t *testing.T) {
	suite.Run(t, new(RefundTestSuite))
}