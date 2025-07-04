package logger

import (
	"context"
	mdwlogger "sc/internal/middleware/logger"

	"go.uber.org/zap"
)

type ContextLogger struct {
	ctx context.Context
}

func FromContext(ctx context.Context) *ContextLogger {
	return &ContextLogger{ctx: ctx}
}

func (c *ContextLogger) Info(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("request_id", mdwlogger.GetRequestID(c.ctx)))
	Log.Info(msg, fields...)
}

func (c *ContextLogger) Error(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("request_id", mdwlogger.GetRequestID(c.ctx)))
	Log.Error(msg, fields...)
}

func (c *ContextLogger) Warn(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("request_id", mdwlogger.GetRequestID(c.ctx)))
	Log.Warn(msg, fields...)
}

func (c *ContextLogger) Debug(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("request_id", mdwlogger.GetRequestID(c.ctx)))
	Log.Debug(msg, fields...)
}
