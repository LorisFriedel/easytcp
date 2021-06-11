// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DarthPestilane/easytcp/packet (interfaces: Packer)

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	packet "github.com/DarthPestilane/easytcp/packet"
	gomock "github.com/golang/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// Pack mocks base method.
func (m *MockPacker) Pack(arg0 packet.Message) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pack", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pack indicates an expected call of Pack.
func (mr *MockPackerMockRecorder) Pack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pack", reflect.TypeOf((*MockPacker)(nil).Pack), arg0)
}

// Unpack mocks base method.
func (m *MockPacker) Unpack(arg0 io.Reader) (packet.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack", arg0)
	ret0, _ := ret[0].(packet.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unpack indicates an expected call of Unpack.
func (mr *MockPackerMockRecorder) Unpack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockPacker)(nil).Unpack), arg0)
}
