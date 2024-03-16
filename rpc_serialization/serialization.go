package rpc_serialization

type Serialization interface {
	Marshal(interface{}) ([]byte, error)
	UnMarshal([]byte, interface{}) error
}

const (
	Proto   = "proto"   // protobuf
	MsgPack = "msgpack" // msgpack
	Json    = "json"    // json
)

func init() {
	registerSerialization("proto", DefaultSerialization)
}

func registerSerialization(name string, serialization Serialization) {
	if serializationMap == nil {
		serializationMap = make(map[string]Serialization)
	}
	serializationMap[name] = serialization
}

// GetSerialization get a Serialization by a serialization name
func GetSerialization(name string) Serialization {
	if v, ok := serializationMap[name]; ok {
		return v
	}
	return DefaultSerialization
}

var serializationMap = make(map[string]Serialization)

var DefaultSerialization = NewPbSerialization()
