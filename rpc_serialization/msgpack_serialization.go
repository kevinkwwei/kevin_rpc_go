package rpc_serialization

import (
	"bytes"
	"errors"
	"github.com/vmihailenco/msgpack"
)

type MsgpackSerialization struct{}

func (m_s *MsgpackSerialization) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("marshal nil interface{}")
	}

	var buf bytes.Buffer
	encoder := msgpack.NewEncoder(&buf)
	err := encoder.Encode(v)
	return buf.Bytes(), err
}

func (m_s *MsgpackSerialization) UnMarshal(data []byte, v interface{}) error {
	if data == nil || len(data) == 0 {
		return errors.New("unmarshal nil or empty bytes")
	}

	decoder := msgpack.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(v)
	return err
}

func NewPbSerialization() Serialization {
	return &MsgpackSerialization{}
}
