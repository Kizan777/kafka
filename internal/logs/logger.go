package logs

import (
	"context"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugaredLogger *zap.SugaredLogger
	defLevel      = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func NewLogger(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defLevel
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(cfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(w), level)
	return zap.New(core, options...).Sugar()
}

type ctxLoggerKey struct{}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	logger := ctx.Value(ctxLoggerKey{})
	if logger != nil {
		return logger.(*zap.SugaredLogger)
	}
	return sugaredLogger
}
