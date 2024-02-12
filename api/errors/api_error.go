package errors

import (
	"net/http"
)

type ApiError struct {
	Message       string
	StatusCode    int
	InternalError error
}

func (s *ApiError) Error() string {
	return s.Message
}

func GetStatusCodeFromError(err error) int {

	if apiError, ok := err.(*ApiError); ok {
		return apiError.StatusCode
	}

	return http.StatusInternalServerError
}
