package codec

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

var _ Codec = &Protobuf{}
var DefaultProtobuf = &Protobuf{}

type Protobuf struct{}

func (p Protobuf) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("invalid data type")
	}
	return proto.Marshal(msg)
}

func (p Protobuf) Unmarshal(b []byte, data interface{}) error {
	msg, ok := data.(proto.Message)
	if !ok {
		return fmt.Errorf("invalid data type")
	}
	err := proto.Unmarshal(b, msg)
	return err
}
