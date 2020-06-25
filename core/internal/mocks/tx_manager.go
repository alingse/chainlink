// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	accounts "github.com/ethereum/go-ethereum/accounts"
	assets "github.com/smartcontractkit/chainlink/core/assets"

	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	eth "github.com/smartcontractkit/chainlink/core/services/eth"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	null "gopkg.in/guregu/null.v3"

	store "github.com/smartcontractkit/chainlink/core/store"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxManager is an autogenerated mock type for the TxManager type
type TxManager struct {
	mock.Mock
}

// BumpGasUntilSafe provides a mock function with given fields: hash
func (_m *TxManager) BumpGasUntilSafe(hash common.Hash) (*models.TxReceipt, store.AttemptState, error) {
	ret := _m.Called(hash)

	var r0 *models.TxReceipt
	if rf, ok := ret.Get(0).(func(common.Hash) *models.TxReceipt); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TxReceipt)
		}
	}

	var r1 store.AttemptState
	if rf, ok := ret.Get(1).(func(common.Hash) store.AttemptState); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Get(1).(store.AttemptState)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(common.Hash) error); ok {
		r2 = rf(hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Call provides a mock function with given fields: result, method, args
func (_m *TxManager) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckAttempt provides a mock function with given fields: txAttempt, blockHeight
func (_m *TxManager) CheckAttempt(txAttempt *models.TxAttempt, blockHeight uint64) (*models.TxReceipt, store.AttemptState, error) {
	ret := _m.Called(txAttempt, blockHeight)

	var r0 *models.TxReceipt
	if rf, ok := ret.Get(0).(func(*models.TxAttempt, uint64) *models.TxReceipt); ok {
		r0 = rf(txAttempt, blockHeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TxReceipt)
		}
	}

	var r1 store.AttemptState
	if rf, ok := ret.Get(1).(func(*models.TxAttempt, uint64) store.AttemptState); ok {
		r1 = rf(txAttempt, blockHeight)
	} else {
		r1 = ret.Get(1).(store.AttemptState)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*models.TxAttempt, uint64) error); ok {
		r2 = rf(txAttempt, blockHeight)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Connect provides a mock function with given fields: _a0
func (_m *TxManager) Connect(_a0 *models.Head) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Head) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connected provides a mock function with given fields:
func (_m *TxManager) Connected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ContractLINKBalance provides a mock function with given fields: wr
func (_m *TxManager) ContractLINKBalance(wr models.WithdrawalRequest) (assets.Link, error) {
	ret := _m.Called(wr)

	var r0 assets.Link
	if rf, ok := ret.Get(0).(func(models.WithdrawalRequest) assets.Link); ok {
		r0 = rf(wr)
	} else {
		r0 = ret.Get(0).(assets.Link)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.WithdrawalRequest) error); ok {
		r1 = rf(wr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTx provides a mock function with given fields: to, data
func (_m *TxManager) CreateTx(to common.Address, data []byte) (*models.Tx, error) {
	ret := _m.Called(to, data)

	var r0 *models.Tx
	if rf, ok := ret.Get(0).(func(common.Address, []byte) *models.Tx); ok {
		r0 = rf(to, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, []byte) error); ok {
		r1 = rf(to, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTxWithEth provides a mock function with given fields: from, to, value
func (_m *TxManager) CreateTxWithEth(from common.Address, to common.Address, value *assets.Eth) (*models.Tx, error) {
	ret := _m.Called(from, to, value)

	var r0 *models.Tx
	if rf, ok := ret.Get(0).(func(common.Address, common.Address, *assets.Eth) *models.Tx); ok {
		r0 = rf(from, to, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address, *assets.Eth) error); ok {
		r1 = rf(from, to, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTxWithGas provides a mock function with given fields: surrogateID, to, data, gasPriceWei, gasLimit
func (_m *TxManager) CreateTxWithGas(surrogateID null.String, to common.Address, data []byte, gasPriceWei *big.Int, gasLimit uint64) (*models.Tx, error) {
	ret := _m.Called(surrogateID, to, data, gasPriceWei, gasLimit)

	var r0 *models.Tx
	if rf, ok := ret.Get(0).(func(null.String, common.Address, []byte, *big.Int, uint64) *models.Tx); ok {
		r0 = rf(surrogateID, to, data, gasPriceWei, gasLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(null.String, common.Address, []byte, *big.Int, uint64) error); ok {
		r1 = rf(surrogateID, to, data, gasPriceWei, gasLimit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disconnect provides a mock function with given fields:
func (_m *TxManager) Disconnect() {
	_m.Called()
}

// GetBlockByNumber provides a mock function with given fields: hex
func (_m *TxManager) GetBlockByNumber(hex string) (models.Block, error) {
	ret := _m.Called(hex)

	var r0 models.Block
	if rf, ok := ret.Get(0).(func(string) models.Block); ok {
		r0 = rf(hex)
	} else {
		r0 = ret.Get(0).(models.Block)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeight provides a mock function with given fields:
func (_m *TxManager) GetBlockHeight() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChainID provides a mock function with given fields:
func (_m *TxManager) GetChainID() (*big.Int, error) {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetERC20Balance provides a mock function with given fields: address, contractAddress
func (_m *TxManager) GetERC20Balance(address common.Address, contractAddress common.Address) (*big.Int, error) {
	ret := _m.Called(address, contractAddress)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *big.Int); ok {
		r0 = rf(address, contractAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(address, contractAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEthBalance provides a mock function with given fields: address
func (_m *TxManager) GetEthBalance(address common.Address) (*assets.Eth, error) {
	ret := _m.Called(address)

	var r0 *assets.Eth
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Eth); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Eth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLINKBalance provides a mock function with given fields: address
func (_m *TxManager) GetLINKBalance(address common.Address) (*assets.Link, error) {
	ret := _m.Called(address)

	var r0 *assets.Link
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Link); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlock provides a mock function with given fields:
func (_m *TxManager) GetLatestBlock() (models.Block, error) {
	ret := _m.Called()

	var r0 models.Block
	if rf, ok := ret.Get(0).(func() models.Block); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Block)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: q
func (_m *TxManager) GetLogs(q ethereum.FilterQuery) ([]models.Log, error) {
	ret := _m.Called(q)

	var r0 []models.Log
	if rf, ok := ret.Get(0).(func(ethereum.FilterQuery) []models.Log); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ethereum.FilterQuery) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: address
func (_m *TxManager) GetNonce(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxReceipt provides a mock function with given fields: hash
func (_m *TxManager) GetTxReceipt(hash common.Hash) (*models.TxReceipt, error) {
	ret := _m.Called(hash)

	var r0 *models.TxReceipt
	if rf, ok := ret.Get(0).(func(common.Hash) *models.TxReceipt); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TxReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GethClient provides a mock function with given fields: _a0
func (_m *TxManager) GethClient(_a0 func(eth.GethClient) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(eth.GethClient) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextActiveAccount provides a mock function with given fields:
func (_m *TxManager) NextActiveAccount() *store.ManagedAccount {
	ret := _m.Called()

	var r0 *store.ManagedAccount
	if rf, ok := ret.Get(0).(func() *store.ManagedAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.ManagedAccount)
		}
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: head
func (_m *TxManager) OnNewLongestChain(head models.Head) {
	_m.Called(head)
}

// Register provides a mock function with given fields: _a0
func (_m *TxManager) Register(_a0 []accounts.Account) {
	_m.Called(_a0)
}

// SendRawTx provides a mock function with given fields: bytes
func (_m *TxManager) SendRawTx(bytes []byte) (common.Hash, error) {
	ret := _m.Called(bytes)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func([]byte) common.Hash); ok {
		r0 = rf(bytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(bytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignedRawTxWithBumpedGas provides a mock function with given fields: originalTx, gasLimit, gasPrice
func (_m *TxManager) SignedRawTxWithBumpedGas(originalTx models.Tx, gasLimit uint64, gasPrice big.Int) ([]byte, error) {
	ret := _m.Called(originalTx, gasLimit, gasPrice)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(models.Tx, uint64, big.Int) []byte); ok {
		r0 = rf(originalTx, gasLimit, gasPrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Tx, uint64, big.Int) error); ok {
		r1 = rf(originalTx, gasLimit, gasPrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: _a0, _a1, _a2
func (_m *TxManager) Subscribe(_a0 context.Context, _a1 interface{}, _a2 ...interface{}) (eth.Subscription, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) eth.Subscription); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToLogs provides a mock function with given fields: ctx, channel, q
func (_m *TxManager) SubscribeToLogs(ctx context.Context, channel chan<- models.Log, q ethereum.FilterQuery) (eth.Subscription, error) {
	ret := _m.Called(ctx, channel, q)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- models.Log, ethereum.FilterQuery) eth.Subscription); ok {
		r0 = rf(ctx, channel, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- models.Log, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, channel, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToNewHeads provides a mock function with given fields: ctx, channel
func (_m *TxManager) SubscribeToNewHeads(ctx context.Context, channel chan<- types.Header) (eth.Subscription, error) {
	ret := _m.Called(ctx, channel)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- types.Header) eth.Subscription); ok {
		r0 = rf(ctx, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- types.Header) error); ok {
		r1 = rf(ctx, channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
