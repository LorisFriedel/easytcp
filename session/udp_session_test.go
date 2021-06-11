package session

import (
	"fmt"
	"github.com/DarthPestilane/easytcp/packet"
	"github.com/DarthPestilane/easytcp/packet/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestNewUDP(t *testing.T) {
	var sess Session
	assert.NotPanics(t, func() {
		sess = NewUDP(nil, nil, nil, nil)
	})
	assert.NotNil(t, sess)
	s, ok := sess.(*UDPSession)
	assert.True(t, ok)
	assert.NotNil(t, s.closed)
	assert.NotNil(t, s.ackQueue)
	assert.NotNil(t, s.reqQueue)
	assert.NotNil(t, s.log)
}

func TestUDPSession_Close(t *testing.T) {
	sess := NewUDP(nil, nil, nil, nil)
	sess.Close()
	var ok bool
	_, ok = <-sess.closed
	assert.False(t, ok)
	_, ok = <-sess.reqQueue
	assert.False(t, ok)
	_, ok = <-sess.ackQueue
	assert.False(t, ok)
}

func TestUDPSession_ID(t *testing.T) {
	sess := NewUDP(nil, nil, nil, nil)
	assert.NotEmpty(t, sess.ID())
	assert.Equal(t, sess.id, sess.ID())
}

func TestUDPSession_MsgCodec(t *testing.T) {
	sess := NewUDP(nil, nil, nil, &packet.JsonCodec{})
	assert.NotNil(t, sess.MsgCodec())
	assert.Equal(t, sess.msgCodec, &packet.JsonCodec{})
	assert.Equal(t, sess.msgCodec, sess.MsgCodec())
}

func TestUDPSession_ReadIncomingMsg(t *testing.T) {
	t.Run("when packer unpack failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Unpack(gomock.Any()).Return(nil, fmt.Errorf("some err"))

		sess := NewUDP(nil, nil, packer, nil)
		go func() { <-sess.reqQueue }()
		assert.Error(t, sess.ReadIncomingMsg([]byte("test")))
		sess.Close()
	})
	t.Run("when reqQueue closed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := mock.NewMockMessage(ctrl)

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Unpack(gomock.Any()).Return(msg, nil)

		sess := NewUDP(nil, nil, packer, nil)
		sess.Close() // close first
		assert.NoError(t, sess.ReadIncomingMsg([]byte("test")))
	})
	t.Run("when succeed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := mock.NewMockMessage(ctrl)

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Unpack(gomock.Any()).Return(msg, nil)

		sess := NewUDP(nil, nil, packer, nil)
		go func() { <-sess.reqQueue }()
		assert.NoError(t, sess.ReadIncomingMsg([]byte("test")))
		sess.Close()
	})
}

func TestUDPSession_RecvReq(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	msg := mock.NewMockMessage(ctrl)

	sess := NewUDP(nil, nil, nil, nil)
	go func() { sess.reqQueue <- msg }()
	req, ok := <-sess.RecvReq()
	assert.True(t, ok)
	assert.Equal(t, req, msg)
	sess.Close()
	_, ok = <-sess.RecvReq()
	assert.False(t, ok)
}

func TestUDPSession_SendResp(t *testing.T) {
	t.Run("when pack msg failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		message := mock.NewMockMessage(ctrl)
		codec := mock.NewMockCodec(ctrl)

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Pack(gomock.Any()).Return(nil, fmt.Errorf("some err"))

		sess := NewUDP(nil, nil, packer, codec)
		closed, err := sess.SendResp(message)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "pack response data err")
		assert.False(t, closed)
	})
	t.Run("when ackQueue closed (session's closed)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		message := mock.NewMockMessage(ctrl)
		codec := mock.NewMockCodec(ctrl)

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Pack(gomock.Any()).Return([]byte("test"), nil)

		sess := NewUDP(nil, nil, packer, codec)
		sess.Close()
		closed, err := sess.SendResp(message)
		assert.NoError(t, err)
		assert.True(t, closed)
	})
	t.Run("when succeed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		message := mock.NewMockMessage(ctrl)
		codec := mock.NewMockCodec(ctrl)

		packer := mock.NewMockPacker(ctrl)
		packer.EXPECT().Pack(gomock.Any()).Return([]byte("test"), nil)

		sess := NewUDP(nil, nil, packer, codec)
		go func() { <-sess.ackQueue }()
		closed, err := sess.SendResp(message)
		assert.NoError(t, err)
		assert.False(t, closed)
	})
}

func TestUDPSession_Write(t *testing.T) {
	t.Run("when conn write failed", func(t *testing.T) {
		addr, err := net.ResolveUDPAddr("udp", "localhost:0")
		assert.NoError(t, err)
		conn, err := net.ListenUDP("udp", addr)
		assert.NoError(t, err)

		sess := NewUDP(conn, nil, nil, nil)
		done := make(chan struct{})
		go func() { sess.ackQueue <- []byte("test") }()
		sess.Write(done)
	})
	t.Run("when ackQueue closed", func(t *testing.T) {
		sess := NewUDP(nil, nil, nil, nil)
		done := make(chan struct{})
		go func() { close(done) }()
		sess.Write(done)
	})
	t.Run("when done closed", func(t *testing.T) {
		sess := NewUDP(nil, nil, nil, nil)
		done := make(chan struct{})
		go func() { close(sess.ackQueue) }()
		sess.Write(done)
	})
}
