package accesslog

import (
	"net/http"
	"time"

	"github.com/dd3v/cloud.snippets.ninja/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/access"
)

func Handler(logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		start := time.Now()
		rw := &access.LogResponseWriter{ResponseWriter: c.Response, Status: http.StatusOK}
		c.Response = rw
		// associate request ID request context
		ctx := c.Request.Context()
		ctx = log.WithRequest(ctx, c.Request)
		c.Request = c.Request.WithContext(ctx)
		err := c.Next()
		logger.With(ctx, "duration", time.Now().Sub(start).Milliseconds(), "status", rw.Status).
			Infof("%s %s %s %d %d", c.Request.Method, c.Request.URL.Path, c.Request.Proto, rw.Status, rw.BytesWritten)
		return err
	}
}
