package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	// bikin meta tapi tipedata nya meta
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// meta memilih struct karena tipe datanya tetap
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:status`
}

//menggunakan func kapital artinya sifatnya global
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
