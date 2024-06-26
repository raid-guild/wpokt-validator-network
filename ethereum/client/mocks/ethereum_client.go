// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	client "github.com/dan13ram/wpokt-oracle/ethereum/client"
	mock "github.com/stretchr/testify/mock"

	models "github.com/dan13ram/wpokt-oracle/models"

	types "github.com/ethereum/go-ethereum/core/types"
)

// MockEthereumClient is an autogenerated mock type for the EthereumClient type
type MockEthereumClient struct {
	mock.Mock
}

type MockEthereumClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEthereumClient) EXPECT() *MockEthereumClient_Expecter {
	return &MockEthereumClient_Expecter{mock: &_m.Mock}
}

// Chain provides a mock function with given fields:
func (_m *MockEthereumClient) Chain() models.Chain {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chain")
	}

	var r0 models.Chain
	if rf, ok := ret.Get(0).(func() models.Chain); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Chain)
	}

	return r0
}

// MockEthereumClient_Chain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chain'
type MockEthereumClient_Chain_Call struct {
	*mock.Call
}

// Chain is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) Chain() *MockEthereumClient_Chain_Call {
	return &MockEthereumClient_Chain_Call{Call: _e.mock.On("Chain")}
}

func (_c *MockEthereumClient_Chain_Call) Run(run func()) *MockEthereumClient_Chain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_Chain_Call) Return(_a0 models.Chain) *MockEthereumClient_Chain_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEthereumClient_Chain_Call) RunAndReturn(run func() models.Chain) *MockEthereumClient_Chain_Call {
	_c.Call.Return(run)
	return _c
}

// Confirmations provides a mock function with given fields:
func (_m *MockEthereumClient) Confirmations() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Confirmations")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// MockEthereumClient_Confirmations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmations'
type MockEthereumClient_Confirmations_Call struct {
	*mock.Call
}

// Confirmations is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) Confirmations() *MockEthereumClient_Confirmations_Call {
	return &MockEthereumClient_Confirmations_Call{Call: _e.mock.On("Confirmations")}
}

func (_c *MockEthereumClient_Confirmations_Call) Run(run func()) *MockEthereumClient_Confirmations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_Confirmations_Call) Return(_a0 uint64) *MockEthereumClient_Confirmations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEthereumClient_Confirmations_Call) RunAndReturn(run func() uint64) *MockEthereumClient_Confirmations_Call {
	_c.Call.Return(run)
	return _c
}

// GetBlockHeight provides a mock function with given fields:
func (_m *MockEthereumClient) GetBlockHeight() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBlockHeight")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEthereumClient_GetBlockHeight_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockHeight'
type MockEthereumClient_GetBlockHeight_Call struct {
	*mock.Call
}

// GetBlockHeight is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) GetBlockHeight() *MockEthereumClient_GetBlockHeight_Call {
	return &MockEthereumClient_GetBlockHeight_Call{Call: _e.mock.On("GetBlockHeight")}
}

func (_c *MockEthereumClient_GetBlockHeight_Call) Run(run func()) *MockEthereumClient_GetBlockHeight_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_GetBlockHeight_Call) Return(_a0 uint64, _a1 error) *MockEthereumClient_GetBlockHeight_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEthereumClient_GetBlockHeight_Call) RunAndReturn(run func() (uint64, error)) *MockEthereumClient_GetBlockHeight_Call {
	_c.Call.Return(run)
	return _c
}

// GetChainID provides a mock function with given fields:
func (_m *MockEthereumClient) GetChainID() (*big.Int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetChainID")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func() (*big.Int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEthereumClient_GetChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChainID'
type MockEthereumClient_GetChainID_Call struct {
	*mock.Call
}

// GetChainID is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) GetChainID() *MockEthereumClient_GetChainID_Call {
	return &MockEthereumClient_GetChainID_Call{Call: _e.mock.On("GetChainID")}
}

func (_c *MockEthereumClient_GetChainID_Call) Run(run func()) *MockEthereumClient_GetChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_GetChainID_Call) Return(_a0 *big.Int, _a1 error) *MockEthereumClient_GetChainID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEthereumClient_GetChainID_Call) RunAndReturn(run func() (*big.Int, error)) *MockEthereumClient_GetChainID_Call {
	_c.Call.Return(run)
	return _c
}

// GetClient provides a mock function with given fields:
func (_m *MockEthereumClient) GetClient() client.EthHTTPClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetClient")
	}

	var r0 client.EthHTTPClient
	if rf, ok := ret.Get(0).(func() client.EthHTTPClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.EthHTTPClient)
		}
	}

	return r0
}

// MockEthereumClient_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type MockEthereumClient_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) GetClient() *MockEthereumClient_GetClient_Call {
	return &MockEthereumClient_GetClient_Call{Call: _e.mock.On("GetClient")}
}

func (_c *MockEthereumClient_GetClient_Call) Run(run func()) *MockEthereumClient_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_GetClient_Call) Return(_a0 client.EthHTTPClient) *MockEthereumClient_GetClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEthereumClient_GetClient_Call) RunAndReturn(run func() client.EthHTTPClient) *MockEthereumClient_GetClient_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionByHash provides a mock function with given fields: txHash
func (_m *MockEthereumClient) GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	ret := _m.Called(txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByHash")
	}

	var r0 *types.Transaction
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (*types.Transaction, bool, error)); ok {
		return rf(txHash)
	}
	if rf, ok := ret.Get(0).(func(string) *types.Transaction); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(txHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(txHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockEthereumClient_GetTransactionByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionByHash'
type MockEthereumClient_GetTransactionByHash_Call struct {
	*mock.Call
}

// GetTransactionByHash is a helper method to define mock.On call
//   - txHash string
func (_e *MockEthereumClient_Expecter) GetTransactionByHash(txHash interface{}) *MockEthereumClient_GetTransactionByHash_Call {
	return &MockEthereumClient_GetTransactionByHash_Call{Call: _e.mock.On("GetTransactionByHash", txHash)}
}

func (_c *MockEthereumClient_GetTransactionByHash_Call) Run(run func(txHash string)) *MockEthereumClient_GetTransactionByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEthereumClient_GetTransactionByHash_Call) Return(_a0 *types.Transaction, _a1 bool, _a2 error) *MockEthereumClient_GetTransactionByHash_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockEthereumClient_GetTransactionByHash_Call) RunAndReturn(run func(string) (*types.Transaction, bool, error)) *MockEthereumClient_GetTransactionByHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionReceipt provides a mock function with given fields: txHash
func (_m *MockEthereumClient) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	ret := _m.Called(txHash)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionReceipt")
	}

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*types.Receipt, error)); ok {
		return rf(txHash)
	}
	if rf, ok := ret.Get(0).(func(string) *types.Receipt); ok {
		r0 = rf(txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEthereumClient_GetTransactionReceipt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionReceipt'
type MockEthereumClient_GetTransactionReceipt_Call struct {
	*mock.Call
}

// GetTransactionReceipt is a helper method to define mock.On call
//   - txHash string
func (_e *MockEthereumClient_Expecter) GetTransactionReceipt(txHash interface{}) *MockEthereumClient_GetTransactionReceipt_Call {
	return &MockEthereumClient_GetTransactionReceipt_Call{Call: _e.mock.On("GetTransactionReceipt", txHash)}
}

func (_c *MockEthereumClient_GetTransactionReceipt_Call) Run(run func(txHash string)) *MockEthereumClient_GetTransactionReceipt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockEthereumClient_GetTransactionReceipt_Call) Return(_a0 *types.Receipt, _a1 error) *MockEthereumClient_GetTransactionReceipt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEthereumClient_GetTransactionReceipt_Call) RunAndReturn(run func(string) (*types.Receipt, error)) *MockEthereumClient_GetTransactionReceipt_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateNetwork provides a mock function with given fields:
func (_m *MockEthereumClient) ValidateNetwork() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidateNetwork")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEthereumClient_ValidateNetwork_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateNetwork'
type MockEthereumClient_ValidateNetwork_Call struct {
	*mock.Call
}

// ValidateNetwork is a helper method to define mock.On call
func (_e *MockEthereumClient_Expecter) ValidateNetwork() *MockEthereumClient_ValidateNetwork_Call {
	return &MockEthereumClient_ValidateNetwork_Call{Call: _e.mock.On("ValidateNetwork")}
}

func (_c *MockEthereumClient_ValidateNetwork_Call) Run(run func()) *MockEthereumClient_ValidateNetwork_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEthereumClient_ValidateNetwork_Call) Return(_a0 error) *MockEthereumClient_ValidateNetwork_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEthereumClient_ValidateNetwork_Call) RunAndReturn(run func() error) *MockEthereumClient_ValidateNetwork_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEthereumClient creates a new instance of MockEthereumClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEthereumClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEthereumClient {
	mock := &MockEthereumClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
