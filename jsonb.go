package pgor

import "encoding/json"

type Json map[string]interface{}

func ToJson[T any](data []byte) T {
	var (
		t T
	)

	json.Unmarshal(data, &t)

	return t
}

func ToBytes(t any) []byte {
	b, _ := json.Marshal(t)

	return b
}
