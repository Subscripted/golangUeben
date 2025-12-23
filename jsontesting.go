package main

import (
	"encoding/json"
)

func dataToJson(data []interface{}) string {
	sData := data[:]
	ret, err := json.Marshal(sData)
	if err != nil {
		return "Fehler beim JSON-Marshalling: " + err.Error()
	}
	return string(ret)
}
