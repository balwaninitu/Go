package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ApiErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type apiErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e apiErr) Message() string {
	return e.ErrMessage
}

func (e apiErr) Status() int {
	return e.ErrStatus
}

func (e apiErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewError(message string, status int, err string, causes []interface{}) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewErrorFromJson(bytes []byte) (ApiErr, error) {
	var apiErr apiErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad request",
	}
}

func NewNotFoundError(message string) ApiErr {
	return apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not found",
	}
}

func NewInternalServerError(message string, err error) ApiErr {
	result := apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal server error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err ApiErr) {
	RespondJson(w, err.Status(), err)
}
