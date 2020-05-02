package errors

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lib/pq"
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
	case *pq.Error:
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" {
				return BadRequest(pgerr.Message)
			}
		}
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
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("")
	}
	return InternalServerError(err.Error())
}
