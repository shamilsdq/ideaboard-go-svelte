package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ParseRequestBody(r *http.Request, dto any) error {
	return json.NewDecoder(r.Body).Decode(dto)
}

func ParseAndValidate(content any, targetDto any) []string {
	jsonBytes, jsonErr := json.Marshal(content)
	if jsonErr != nil {
		return []string{jsonErr.Error()}
	}

	if parseErr := json.Unmarshal(jsonBytes, targetDto); parseErr != nil {
		return []string{parseErr.Error()}
	}

	if validationErr := validator.New().Struct(targetDto); validationErr != nil {
		errors := make([]string, 0)
		for _, err := range validationErr.(validator.ValidationErrors) {
			errors = append(errors, generateErrorString(err))
		}
		return errors
	}

	return nil
}

func generateErrorString(err validator.FieldError) string {
	return ""
}
