package client

import (
	"fmt"
	"net/http"
)

type KontenbaseError struct {
	Message    string `json:"message"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

func (err *KontenbaseError) Error() string {
	return fmt.Sprintf("%d %s %s", err.Status, err.StatusText, err.Message)
}

func internalError(message string) *KontenbaseError {
	return &KontenbaseError{
		Message:    message,
		Status:     http.StatusInternalServerError,
		StatusText: "Internal Server Error",
	}
}
