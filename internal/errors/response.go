package errors

import (
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation"
)

//ErrorResponse - respresents error response
type ErrorResponse struct {
	Status  int         `json:"status" `
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

//StatusCode - returns http status code
func (r ErrorResponse) StatusCode() int {
	return r.Status
}

//Error - returns error message
func (r ErrorResponse) Error() string {
	return r.Message
}

//GetDetails - returns detail infromatopn about error
func (r ErrorResponse) GetDetails() interface{} {
	return r.Details
}

//NotFound - HTTP 404
func NotFound(message string) ErrorResponse {
	if message == "" {
		message = "Not found"
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

//Forbidden - HTTP 403
func Forbidden(message string) ErrorResponse {
	if message == "" {
		message = "Forbidden"
	}
	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: message,
	}
}

//BadRequest - HTTP 400
func BadRequest(message string) ErrorResponse {
	if message == "" {
		message = "Bad request"
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

//Unauthorized - HTTP 401
func Unauthorized(message string) ErrorResponse {
	if message == "" {
		message = "Unauthorized"
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

//InternalServerError - HTTP 500
func InternalServerError(message string) ErrorResponse {
	if message == "" {
		message = "Internal server error"
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

//GenerateValidationError - ...
func GenerateValidationError(errs validation.Errors) ErrorResponse {
	var details []invalidField
	var fields []string
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		details = append(details, invalidField{
			Field: field,
			Error: errs[field].Error(),
		})
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "Validation error",
		Details: details,
	}
}
