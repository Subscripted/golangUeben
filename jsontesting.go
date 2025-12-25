package main

import (
	"encoding/json"
)

func dataToJson(v any) ([]byte, error) {
	return json.Marshal(v)
}
