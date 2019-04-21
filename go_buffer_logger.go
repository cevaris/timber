package timber

import (
	"bufio"
	"context"
	"log"
	"strings"
)

// NewGoBufferLogger returns new logger
func NewGoBufferLogger(buf *bufio.Writer) Logger {
	logger := log.New(buf, "", log.LstdFlags|log.LUTC)
	return &goBufferLogger{logger: logger}
}

type goBufferLogger struct {
	logger *log.Logger
}

func (l *goBufferLogger) Info(ctx context.Context, m ...interface{}) {
	if Level == Error || Level == Info {
		l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
	}
}

func (l *goBufferLogger) Error(ctx context.Context, m ...interface{}) {
	l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
}

func (l *goBufferLogger) Debug(ctx context.Context, m ...interface{}) {
	if Level == Debug {
		l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
	}
}
