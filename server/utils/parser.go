package utils

import (
	"encoding/json"
	"net/http"
)

func ParseRequestBody(r *http.Request, dto any) error {
	return json.NewDecoder(r.Body).Decode(dto)
}

func ParseToDto(content interface{}, dto interface{}) error {
	jsonBytes, jsonErr := json.Marshal(content)
	if jsonErr != nil {
		return jsonErr
	}

	if parseErr := json.Unmarshal(jsonBytes, dto); parseErr != nil {
		return parseErr
	}

	return nil
}
