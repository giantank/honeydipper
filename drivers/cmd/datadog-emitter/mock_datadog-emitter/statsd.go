// Copyright 2022 PayPal Inc.

// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this file,
// you can obtain one at https://mit-license.org/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: drivers/cmd/datadog-emitter/statsd.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	statsd "github.com/DataDog/datadog-go/statsd"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockvirtualStatsd is a mock of virtualStatsd interface
type MockvirtualStatsd struct {
	ctrl     *gomock.Controller
	recorder *MockvirtualStatsdMockRecorder
}

// MockvirtualStatsdMockRecorder is the mock recorder for MockvirtualStatsd
type MockvirtualStatsdMockRecorder struct {
	mock *MockvirtualStatsd
}

// NewMockvirtualStatsd creates a new mock instance
func NewMockvirtualStatsd(ctrl *gomock.Controller) *MockvirtualStatsd {
	mock := &MockvirtualStatsd{ctrl: ctrl}
	mock.recorder = &MockvirtualStatsdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockvirtualStatsd) EXPECT() *MockvirtualStatsdMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockvirtualStatsd) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockvirtualStatsdMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockvirtualStatsd)(nil).Close))
}

// Event mocks base method
func (m *MockvirtualStatsd) Event(arg0 *statsd.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Event", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Event indicates an expected call of Event
func (mr *MockvirtualStatsdMockRecorder) Event(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockvirtualStatsd)(nil).Event), arg0)
}

// Incr mocks base method
func (m *MockvirtualStatsd) Incr(arg0 string, arg1 []string, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Incr indicates an expected call of Incr
func (mr *MockvirtualStatsdMockRecorder) Incr(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockvirtualStatsd)(nil).Incr), arg0, arg1, arg2)
}

// Gauge mocks base method
func (m *MockvirtualStatsd) Gauge(arg0 string, arg1 float64, arg2 []string, arg3 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gauge", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gauge indicates an expected call of Gauge
func (mr *MockvirtualStatsdMockRecorder) Gauge(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockvirtualStatsd)(nil).Gauge), arg0, arg1, arg2, arg3)
}