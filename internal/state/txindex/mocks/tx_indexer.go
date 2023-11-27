// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/cometbft/cometbft/libs/log"
	mock "github.com/stretchr/testify/mock"

	query "github.com/cometbft/cometbft/internal/pubsub/query"

	txindex "github.com/cometbft/cometbft/internal/state/txindex"

	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

// TxIndexer is an autogenerated mock type for the TxIndexer type
type TxIndexer struct {
	mock.Mock
}

// AddBatch provides a mock function with given fields: b
func (_m *TxIndexer) AddBatch(b *txindex.Batch) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txindex.Batch) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: hash
func (_m *TxIndexer) Get(hash []byte) (*v1.TxResult, error) {
	ret := _m.Called(hash)

	var r0 *v1.TxResult
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*v1.TxResult, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func([]byte) *v1.TxResult); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.TxResult)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetainHeight provides a mock function with given fields:
func (_m *TxIndexer) GetRetainHeight() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index provides a mock function with given fields: result
func (_m *TxIndexer) Index(result *v1.TxResult) error {
	ret := _m.Called(result)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.TxResult) error); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Prune provides a mock function with given fields: retainHeight
func (_m *TxIndexer) Prune(retainHeight int64) (int64, int64, error) {
	ret := _m.Called(retainHeight)

	var r0 int64
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int64) (int64, int64, error)); ok {
		return rf(retainHeight)
	}
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(retainHeight)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64) int64); ok {
		r1 = rf(retainHeight)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int64) error); ok {
		r2 = rf(retainHeight)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Search provides a mock function with given fields: ctx, q
func (_m *TxIndexer) Search(ctx context.Context, q *query.Query) ([]*v1.TxResult, error) {
	ret := _m.Called(ctx, q)

	var r0 []*v1.TxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Query) ([]*v1.TxResult, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *query.Query) []*v1.TxResult); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.TxResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *query.Query) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetLogger provides a mock function with given fields: l
func (_m *TxIndexer) SetLogger(l log.Logger) {
	_m.Called(l)
}

// SetRetainHeight provides a mock function with given fields: retainHeight
func (_m *TxIndexer) SetRetainHeight(retainHeight int64) error {
	ret := _m.Called(retainHeight)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(retainHeight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTxIndexer creates a new instance of TxIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxIndexer(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxIndexer {
	mock := &TxIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
