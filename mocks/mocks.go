// Code generated by MockGen. DO NOT EDIT.
// Source: scanner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRowsScanner is a mock of RowsScanner interface
type MockRowsScanner struct {
	ctrl     *gomock.Controller
	recorder *MockRowsScannerMockRecorder
}

// MockRowsScannerMockRecorder is the mock recorder for MockRowsScanner
type MockRowsScannerMockRecorder struct {
	mock *MockRowsScanner
}

// NewMockRowsScanner creates a new mock instance
func NewMockRowsScanner(ctrl *gomock.Controller) *MockRowsScanner {
	mock := &MockRowsScanner{ctrl: ctrl}
	mock.recorder = &MockRowsScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRowsScanner) EXPECT() *MockRowsScannerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockRowsScanner) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRowsScannerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRowsScanner)(nil).Close))
}

// Scan mocks base method
func (m *MockRowsScanner) Scan(dest ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockRowsScannerMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRowsScanner)(nil).Scan), dest...)
}

// Columns mocks base method
func (m *MockRowsScanner) Columns() ([]string, error) {
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Columns indicates an expected call of Columns
func (mr *MockRowsScannerMockRecorder) Columns() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockRowsScanner)(nil).Columns))
}

// ColumnTypes mocks base method
func (m *MockRowsScanner) ColumnTypes() ([]*sql.ColumnType, error) {
	ret := m.ctrl.Call(m, "ColumnTypes")
	ret0, _ := ret[0].([]*sql.ColumnType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ColumnTypes indicates an expected call of ColumnTypes
func (mr *MockRowsScannerMockRecorder) ColumnTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnTypes", reflect.TypeOf((*MockRowsScanner)(nil).ColumnTypes))
}

// Err mocks base method
func (m *MockRowsScanner) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockRowsScannerMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRowsScanner)(nil).Err))
}

// Next mocks base method
func (m *MockRowsScanner) Next() bool {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockRowsScannerMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRowsScanner)(nil).Next))
}

// NextResultSet mocks base method
func (m *MockRowsScanner) NextResultSet() bool {
	ret := m.ctrl.Call(m, "NextResultSet")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NextResultSet indicates an expected call of NextResultSet
func (mr *MockRowsScannerMockRecorder) NextResultSet() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextResultSet", reflect.TypeOf((*MockRowsScanner)(nil).NextResultSet))
}

// MockScanner is a mock of Scanner interface
type MockScanner struct {
	ctrl     *gomock.Controller
	recorder *MockScannerMockRecorder
}

// MockScannerMockRecorder is the mock recorder for MockScanner
type MockScannerMockRecorder struct {
	mock *MockScanner
}

// NewMockScanner creates a new mock instance
func NewMockScanner(ctrl *gomock.Controller) *MockScanner {
	mock := &MockScanner{ctrl: ctrl}
	mock.recorder = &MockScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScanner) EXPECT() *MockScannerMockRecorder {
	return m.recorder
}

// Scan mocks base method
func (m *MockScanner) Scan(dest ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range dest {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *MockScannerMockRecorder) Scan(dest ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockScanner)(nil).Scan), dest...)
}