package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta
	Data interface{}
}

type Meta struct {
	Message string
	Code    int
	Status  string
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonReponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonReponse
}

func FormatValidationError(err error) []string {

	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors

}
