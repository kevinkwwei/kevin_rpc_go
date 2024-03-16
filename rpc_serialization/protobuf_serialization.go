package rpc_serialization

import (
	"errors"
	"github.com/golang/protobuf/proto"
)

type ProtobufSerialization struct{}

func (pb_s *ProtobufSerialization) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("marshal nil interface{}")
	}
	if pm, ok := v.(proto.Marshaler); ok {
		// 可以 marshal 自身，无需 buffer
		return pm.Marshal()
	}
	buffer := bufferPool.Get().(*cachedBuffer)
	protoMsg := v.(proto.Message)
	lastMarshaledSize := make([]byte, 0, buffer.lastMarshaledSize)
	buffer.SetBuf(lastMarshaledSize)
	buffer.Reset()

	if err := buffer.Marshal(protoMsg); err != nil {
		return nil, err
	}
	data := buffer.Bytes()
	buffer.lastMarshaledSize = upperLimit(len(data))
	buffer.SetBuf(nil)
	bufferPool.Put(buffer)

	return data, nil
}

func (pb_s *ProtobufSerialization) UnMarshal(data []byte, v interface{}) error {
	if data == nil || len(data) == 0 {
		return errors.New("unmarshal nil or empty bytes")
	}

	protoMsg := v.(proto.Message)
	protoMsg.Reset()

	if pu, ok := protoMsg.(proto.Unmarshaler); ok {
		// 可以 unmarshal 自身，无需 buffer
		return pu.Unmarshal(data)
	}

	buffer := bufferPool.Get().(*cachedBuffer)
	buffer.SetBuf(data)
	err := buffer.Unmarshal(protoMsg)
	buffer.SetBuf(nil)
	bufferPool.Put(buffer)
	return err
}

func NewPbSerialization() Serialization {
	return &ProtobufSerialization{}
}
