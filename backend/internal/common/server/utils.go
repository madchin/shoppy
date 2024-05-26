package server

import (
	"encoding/json"
	"io"
)

func DecodeJSON[T any](reader io.ReadCloser) (*T, error) {
	var msg *T
	if err := json.NewDecoder(reader).Decode(&msg); err != nil {
		return nil, err
	}
	return msg, nil
}
