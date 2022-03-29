package helper

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var resErrors []string

	if err != nil {
		var jsErr *json.UnmarshalTypeError
		if errors.As(err, &jsErr) {
			fmt.Println("the json is invalid", err.Error())
			resErrors = append(resErrors, err.Error())
		} else {
			fmt.Println("this is something else")
			for _, e := range err.(validator.ValidationErrors) {
				resErrors = append(resErrors, e.Error())
			}
		}
	}
	return resErrors

}
