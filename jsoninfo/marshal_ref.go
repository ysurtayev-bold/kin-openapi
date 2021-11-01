package jsoninfo

import (
	"encoding/json"
)

func MarshalRef(value string, otherwise interface{}) ([]byte, error) {
	if len(value) > 0 {
		return json.Marshal(&refProps{
			Ref: value,
		})
	}
	return json.Marshal(otherwise)
}

func UnmarshalRef(data []byte, destRef *string, protoNumber *uint64, destOtherwise interface{}) error {
	refProps := &refProps{}
	if err := json.Unmarshal(data, refProps); err == nil {
		ref := refProps.Ref
		if len(ref) > 0 {
			*destRef = ref
			*protoNumber = refProps.ProtoNumber
			return nil
		}
	}
	return json.Unmarshal(data, destOtherwise)
}

type refProps struct {
	Ref string `json:"$ref,omitempty"`

	// Proto-compatible
	ProtoNumber uint64 `json:"protoNumber,omitempty" yaml:"protoNumber,omitempty"`
}
