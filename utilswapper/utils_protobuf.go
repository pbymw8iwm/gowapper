package utilswapper

import (
	"errors"
	"reflect"

	"github.com/golang/protobuf/proto"
)

func ProtobufUnmarshal(proto_name string, data []byte) (interface{}, error) {
	if len(data) < 1 {
		return nil, errors.New("protobuf data too short")
	}
	msg := reflect.New(proto.MessageType(proto_name).Elem()).Interface()
	return msg, proto.UnmarshalMerge(data, msg.(proto.Message))
}

func ProtobufMarshal(msg interface{}) ([]byte, error) {
	data, err := proto.Marshal(msg.(proto.Message))
	return data, err
}
