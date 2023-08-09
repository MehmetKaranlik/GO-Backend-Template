package Validation

import (
	"encoding/json"
	"errors"
	validator2 "github.com/go-playground/validator/v10"
	"io"
)

type CustomJsonDecoder struct {
	*json.Decoder
}

func NewCustomJsonDecoder(reader io.Reader) *CustomJsonDecoder {
	return &CustomJsonDecoder{Decoder: json.NewDecoder(reader)}
}

func (c *CustomJsonDecoder) Decode(v interface{}) error {
	err := c.Decoder.Decode(v)
	if err != nil {
		return err
	}

	validator := validator2.New()
	err = validator.Struct(v)
	if err != nil {
		errText := ""
		for _, err := range err.(validator2.ValidationErrors) {
			errText += err.Error() + ", "
		}
		return errors.New(errText)
	}
	return nil
}
