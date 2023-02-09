// Copyright 2022 PayPal Inc.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this file,
// you can obtain one at https://mit-license.org/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dipper/rpc.go

// Package mock_dipper is a generated GoMock package.
package mock_dipper

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRPCCallerStub is a mock of RPCCallerStub interface.
type MockRPCCallerStub struct {
	ctrl     *gomock.Controller
	recorder *MockRPCCallerStubMockRecorder
}

// MockRPCCallerStubMockRecorder is the mock recorder for MockRPCCallerStub.
type MockRPCCallerStubMockRecorder struct {
	mock *MockRPCCallerStub
}

// NewMockRPCCallerStub creates a new mock instance.
func NewMockRPCCallerStub(ctrl *gomock.Controller) *MockRPCCallerStub {
	mock := &MockRPCCallerStub{ctrl: ctrl}
	mock.recorder = &MockRPCCallerStubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRPCCallerStub) EXPECT() *MockRPCCallerStubMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockRPCCallerStub) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockRPCCallerStubMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockRPCCallerStub)(nil).GetName))
}

// GetStream mocks base method.
func (m *MockRPCCallerStub) GetStream(feature string) io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream", feature)
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStream indicates an expected call of GetStream.
func (mr *MockRPCCallerStubMockRecorder) GetStream(feature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockRPCCallerStub)(nil).GetStream), feature)
}

// MockRPCCaller is a mock of RPCCaller interface.
type MockRPCCaller struct {
	ctrl     *gomock.Controller
	recorder *MockRPCCallerMockRecorder
}

// MockRPCCallerMockRecorder is the mock recorder for MockRPCCaller.
type MockRPCCallerMockRecorder struct {
	mock *MockRPCCaller
}

// NewMockRPCCaller creates a new mock instance.
func NewMockRPCCaller(ctrl *gomock.Controller) *MockRPCCaller {
	mock := &MockRPCCaller{ctrl: ctrl}
	mock.recorder = &MockRPCCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRPCCaller) EXPECT() *MockRPCCallerMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockRPCCaller) Call(feature, method string, params interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", feature, method, params)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockRPCCallerMockRecorder) Call(feature, method, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRPCCaller)(nil).Call), feature, method, params)
}

// CallNoWait mocks base method.
func (m *MockRPCCaller) CallNoWait(feature, method string, params interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallNoWait", feature, method, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallNoWait indicates an expected call of CallNoWait.
func (mr *MockRPCCallerMockRecorder) CallNoWait(feature, method, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallNoWait", reflect.TypeOf((*MockRPCCaller)(nil).CallNoWait), feature, method, params)
}

// CallRaw mocks base method.
func (m *MockRPCCaller) CallRaw(feature, method string, params []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRaw", feature, method, params)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallRaw indicates an expected call of CallRaw.
func (mr *MockRPCCallerMockRecorder) CallRaw(feature, method, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRaw", reflect.TypeOf((*MockRPCCaller)(nil).CallRaw), feature, method, params)
}

// CallRawNoWait mocks base method.
func (m *MockRPCCaller) CallRawNoWait(feature, method string, params []byte, rpcID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallRawNoWait", feature, method, params, rpcID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CallRawNoWait indicates an expected call of CallRawNoWait.
func (mr *MockRPCCallerMockRecorder) CallRawNoWait(feature, method, params, rpcID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallRawNoWait", reflect.TypeOf((*MockRPCCaller)(nil).CallRawNoWait), feature, method, params, rpcID)
}

// GetName mocks base method.
func (m *MockRPCCaller) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockRPCCallerMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockRPCCaller)(nil).GetName))
}
