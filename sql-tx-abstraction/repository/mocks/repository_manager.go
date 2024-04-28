// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository_manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"
	repository "sql-playground/repository"

	gomock "github.com/golang/mock/gomock"
)

// MockDBOps is a mock of DBOps interface.
type MockDBOps struct {
	ctrl     *gomock.Controller
	recorder *MockDBOpsMockRecorder
}

// MockDBOpsMockRecorder is the mock recorder for MockDBOps.
type MockDBOpsMockRecorder struct {
	mock *MockDBOps
}

// NewMockDBOps creates a new mock instance.
func NewMockDBOps(ctrl *gomock.Controller) *MockDBOps {
	mock := &MockDBOps{ctrl: ctrl}
	mock.recorder = &MockDBOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBOps) EXPECT() *MockDBOpsMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockDBOps) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDBOpsMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDBOps)(nil).Exec), varargs...)
}

// Prepare mocks base method.
func (m *MockDBOps) Prepare(query string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", query)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockDBOpsMockRecorder) Prepare(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockDBOps)(nil).Prepare), query)
}

// Query mocks base method.
func (m *MockDBOps) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDBOpsMockRecorder) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDBOps)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockDBOps) QueryRow(query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockDBOpsMockRecorder) QueryRow(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockDBOps)(nil).QueryRow), varargs...)
}

// MockRepositoryManager is a mock of RepositoryManager interface.
type MockRepositoryManager struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryManagerMockRecorder
}

// MockRepositoryManagerMockRecorder is the mock recorder for MockRepositoryManager.
type MockRepositoryManagerMockRecorder struct {
	mock *MockRepositoryManager
}

// NewMockRepositoryManager creates a new mock instance.
func NewMockRepositoryManager(ctrl *gomock.Controller) *MockRepositoryManager {
	mock := &MockRepositoryManager{ctrl: ctrl}
	mock.recorder = &MockRepositoryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryManager) EXPECT() *MockRepositoryManagerMockRecorder {
	return m.recorder
}

// ExampleRepo mocks base method.
func (m *MockRepositoryManager) ExampleRepo() repository.ExampleRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExampleRepo")
	ret0, _ := ret[0].(repository.ExampleRepo)
	return ret0
}

// ExampleRepo indicates an expected call of ExampleRepo.
func (mr *MockRepositoryManagerMockRecorder) ExampleRepo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExampleRepo", reflect.TypeOf((*MockRepositoryManager)(nil).ExampleRepo))
}

// RunAtomic mocks base method.
func (m *MockRepositoryManager) RunAtomic(fn repository.AtomicFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunAtomic", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAtomic indicates an expected call of RunAtomic.
func (mr *MockRepositoryManagerMockRecorder) RunAtomic(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAtomic", reflect.TypeOf((*MockRepositoryManager)(nil).RunAtomic), fn)
}
