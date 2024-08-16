package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// ValidateStruct
// [GIT] https://github.com/go-playground/validator
// example :
//
//	type Name struct {
//		value1      uint   `json:"value1" validate:"required"`
//		value2 		string `json:"value2"`
//	}
func ValidateStruct(dataStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("%s: %s", err.StructField(), err.Tag()))
		}
	} else {
		return nil
	}
	return err
}

func CheckStatus(statusOrder, status string) (res bool) {
	res = true
	if statusOrder != status {
		res = false
	}
	return
}
