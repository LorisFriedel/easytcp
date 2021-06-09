// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DarthPestilane/easytcp/session (interfaces: Session)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	packet "github.com/DarthPestilane/easytcp/packet"
	gomock "github.com/golang/mock/gomock"
)

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSession) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// ID mocks base method.
func (m *MockSession) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSessionMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSession)(nil).ID))
}

// MsgCodec mocks base method.
func (m *MockSession) MsgCodec() packet.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsgCodec")
	ret0, _ := ret[0].(packet.Codec)
	return ret0
}

// MsgCodec indicates an expected call of MsgCodec.
func (mr *MockSessionMockRecorder) MsgCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsgCodec", reflect.TypeOf((*MockSession)(nil).MsgCodec))
}

// RecvReq mocks base method.
func (m *MockSession) RecvReq() <-chan *packet.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvReq")
	ret0, _ := ret[0].(<-chan *packet.Request)
	return ret0
}

// RecvReq indicates an expected call of RecvReq.
func (mr *MockSessionMockRecorder) RecvReq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvReq", reflect.TypeOf((*MockSession)(nil).RecvReq))
}

// SendResp mocks base method.
func (m *MockSession) SendResp(arg0 *packet.Response) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendResp", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendResp indicates an expected call of SendResp.
func (mr *MockSessionMockRecorder) SendResp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResp", reflect.TypeOf((*MockSession)(nil).SendResp), arg0)
}