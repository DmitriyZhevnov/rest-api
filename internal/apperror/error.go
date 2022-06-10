package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound         = NewAppError(nil, "Not found", "", "23243234")
	InternalServerError = NewAppError(nil, "Internal Server error", "", "35345")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func (a *AppError) Error() string {
	return a.Message
}

func (a *AppError) Unwrap() error {
	return a.Err
}

func (a *AppError) Marshal() []byte {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil
	}
	return marshal
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal server error", err.Error(), "0000000")
}
