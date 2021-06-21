package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// SafeUnmarshal decodes JSON into the specifed output object, with
// DisallowUnknownFields enabled.
func SafeUnmarshal(data []byte, output interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buf)
	decoder.DisallowUnknownFields()
	return decoder.Decode(output)
}

// MustMarshal marshals an object to JSON, and panics if any error occurs.
func MustMarshal(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("Error in MustMarshal(): %s", err.Error()))
	}
	return b
}
