package utils

import (
	"encoding/json"
	"net/http"
)

func ParseRequestBody(r *http.Request, dto any) error {
	return json.NewDecoder(r.Body).Decode(dto)
}
