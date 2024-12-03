package mockrepository

import (
	context "context"
	reflect "reflect"

	entity "github.com/gmcriptobox/otus-go-final-project/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIListRepo is a mock of IListRepo interface.
type MockIListRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIListRepoMockRecorder
}

// MockIListRepoMockRecorder is the mock recorder for MockIListRepo.
type MockIListRepoMockRecorder struct {
	mock *MockIListRepo
}

// NewMockIListRepo creates a new mock instance.
func NewMockIListRepo(ctrl *gomock.Controller) *MockIListRepo {
	mock := &MockIListRepo{ctrl: ctrl}
	mock.recorder = &MockIListRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIListRepo) EXPECT() *MockIListRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIListRepo) Add(ctx context.Context, network entity.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, network)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIListRepoMockRecorder) Add(ctx, network interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add",
		reflect.TypeOf((*MockIListRepo)(nil).Add), ctx, network)
}

// GetAll mocks base method.
func (m *MockIListRepo) GetAll(ctx context.Context) ([]entity.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]entity.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIListRepoMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.
		TypeOf((*MockIListRepo)(nil).GetAll), ctx)
}

// IsExists mocks base method.
func (m *MockIListRepo) IsExists(ctx context.Context, ip, mask string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExists", ctx, ip, mask)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExists indicates an expected call of IsExists.
func (mr *MockIListRepoMockRecorder) IsExists(ctx, ip, mask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExists",
		reflect.TypeOf((*MockIListRepo)(nil).IsExists), ctx, ip, mask)
}

// Remove mocks base method.
func (m *MockIListRepo) Remove(ctx context.Context, ip, mask string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, ip, mask)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockIListRepoMockRecorder) Remove(ctx, ip, mask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.
		TypeOf((*MockIListRepo)(nil).Remove), ctx, ip, mask)
}
