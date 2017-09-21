package json

import (
	"encoding/json"
	"net/http"
)

func WriteJson(response http.ResponseWriter, obj interface{}) error {
	json, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(json)
	return nil
}

func ReadJson(data []byte, obj interface{}) {
	_ = json.Unmarshal(data, &obj)
}
