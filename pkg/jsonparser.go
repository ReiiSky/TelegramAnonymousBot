package pkg

import (
	"io"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

/*
	This file containing function that override
	json encoding functionalities from golang standard
	and caused big performance impact.
*/

// JSONUnmarshal method to transform byte to object
func JSONUnmarshal(chunk []byte, object interface{}) error {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(chunk, object)
}

// JSONMarshal transform object to array of byte
func JSONMarshal(object interface{}) ([]byte, error) {
	return jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(object)
}

// UnmarshalByIOReader unmarshal
func UnmarshalByIOReader(reader io.Reader, object interface{}) error {
	chunk, err := ioutil.ReadAll(reader)
	if err == nil {
		err = JSONUnmarshal(chunk, object)
	}
	return err
}
