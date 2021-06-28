package log

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"net/http"
)

const RequestId = "request_id"

type Logger interface {
	// With returns a logger based off the root logger and decorates it with the given context and arguments.
	With(ctx context.Context, args ...interface{}) Logger
	// Debug uses fmt.Sprint to construct and log a message at DEBUG level
	Debug(args ...interface{})
	// Info uses fmt.Sprint to construct and log a message at INFO level
	Info(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at ERROR level
	Error(args ...interface{})
	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debugf(format string, args ...interface{})
	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Infof(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Errorf(format string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

func New(outputPaths []string) Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = outputPaths
	z, _ := config.Build()

	return &logger{z.Sugar()}
}

func NewForTests() (Logger, *observer.ObservedLogs) {
	core, recorded := observer.New(zapcore.InfoLevel)
	z := zap.New(core)
	return &logger{z.Sugar()}, recorded
}

func (l logger) With(ctx context.Context, args ...interface{}) Logger {
	if ctx != nil {
		if id, ok := ctx.Value(RequestId).(string); ok {
			args = append(args, zap.String("request_id", id))
		}
	}
	if len(args) > 0 {
		return &logger{l.SugaredLogger.With(args...)}
	}

	return l
}

func WithRequest(ctx context.Context, request *http.Request) context.Context {
	id := getRequestID(request)
	if id == "" {
		id = uuid.New().String()
	}
	ctx = context.WithValue(ctx, RequestId, id)

	return ctx
}

func getRequestID(req *http.Request) string {
	return req.Header.Get("X-Request-ID")
}
