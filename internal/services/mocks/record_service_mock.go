// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/record.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "git.epam.com/ryan_wang/crud-demo/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRecordService is a mock of RecordService interface.
type MockRecordService struct {
	ctrl     *gomock.Controller
	recorder *MockRecordServiceMockRecorder
}

// MockRecordServiceMockRecorder is the mock recorder for MockRecordService.
type MockRecordServiceMockRecorder struct {
	mock *MockRecordService
}

// NewMockRecordService creates a new mock instance.
func NewMockRecordService(ctrl *gomock.Controller) *MockRecordService {
	mock := &MockRecordService{ctrl: ctrl}
	mock.recorder = &MockRecordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecordService) EXPECT() *MockRecordServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRecordService) Create(ctx context.Context, record *models.Record) (*models.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, record)
	ret0, _ := ret[0].(*models.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRecordServiceMockRecorder) Create(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRecordService)(nil).Create), ctx, record)
}

// Delete mocks base method.
func (m *MockRecordService) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRecordServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecordService)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockRecordService) Get(ctx context.Context, id int64) (*models.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*models.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRecordServiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRecordService)(nil).Get), ctx, id)
}

// Query mocks base method.
func (m *MockRecordService) Query(ctx context.Context, displayName string, pageNum, pageSize int) ([]*models.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, displayName, pageNum, pageSize)
	ret0, _ := ret[0].([]*models.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockRecordServiceMockRecorder) Query(ctx, displayName, pageNum, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRecordService)(nil).Query), ctx, displayName, pageNum, pageSize)
}

// Update mocks base method.
func (m *MockRecordService) Update(ctx context.Context, record *models.Record) (*models.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, record)
	ret0, _ := ret[0].(*models.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRecordServiceMockRecorder) Update(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRecordService)(nil).Update), ctx, record)
}
