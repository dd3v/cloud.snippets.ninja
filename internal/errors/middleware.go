package errors

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/dd3v/cloud.snippets.ninja/internal/rbac"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	validation "github.com/go-ozzo/ozzo-validation"
)

//Handler - error handler HTTP middleware
func Handler() routing.Handler {
	return func(c *routing.Context) (err error) {
		defer func() {
			if err != nil {
				res := buildResponseWithError(err)
				c.Response.WriteHeader(res.StatusCode())
				if err := c.Write(res); err != nil {
					log.Println(err)
				}
				c.Abort()
				err = nil
			}
		}()
		return c.Next()
	}
}

func buildResponseWithError(err error) ErrorResponse {
	switch err.(type) {
	case ErrorResponse:
		return err.(ErrorResponse)
	case validation.Errors:
		return GenerateValidationError(err.(validation.Errors))
	case routing.HTTPError:
		switch err.(routing.HTTPError).StatusCode() {
		case http.StatusNotFound:
			return NotFound("")
		default:
			return ErrorResponse{
				Status:  err.(routing.HTTPError).StatusCode(),
				Message: err.Error(),
			}
		}
	}

	if errors.Is(err, rbac.AccessError) {
		return Forbidden(err.Error())
	}
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("")
	}
	return InternalServerError(err.Error())
}
