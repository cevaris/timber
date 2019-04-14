package timber

import (
	"context"
)

// NewNilLogger returns new logger that does nothing
func NewNilLogger() Logger {
	return &nilLogger{}
}

type nilLogger struct {
}

func (l *nilLogger) Info(ctx context.Context, m ...interface{}) {
}

func (l *nilLogger) Error(ctx context.Context, m ...interface{}) {
}

func (l *nilLogger) Debug(ctx context.Context, m ...interface{}) {
}
