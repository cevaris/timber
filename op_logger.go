package timber

import (
	"context"
	"os"
	"strings"

	"github.com/op/go-logging"
)

// NewOpLogger returns new logger
func NewOpLogger(name string) Logger {
	logger := logging.MustGetLogger(name)
	format := logging.MustStringFormatter(
		// `%{color}%{time:2006-01-02T15:04:05.999Z-07:00.000} %{level:.3s} %{longfunc}%{color:reset} %{message}`,
		`%{color}%{time:2006-01-02T15:04:05.999Z-07:00.000} %{level:.3s} %{color:reset} %{message}`,
	)

	logBackend := logging.NewLogBackend(os.Stdout, "", 0)
	logFormatter := logging.NewBackendFormatter(logBackend, format)
	logging.SetBackend(logFormatter)

	// logging.SetLevel(logging.INFO, "")
	switch Level {
	case Info:
		logging.SetLevel(logging.INFO, "")
		break
	case Error:
		logging.SetLevel(logging.ERROR, "")
		break
	default:
		// defaulting to debug
		logging.SetLevel(logging.DEBUG, "")
	}

	return &opLogger{logger: logger}
}

type opLogger struct {
	logger *logging.Logger
}

func (l *opLogger) Info(ctx context.Context, m ...interface{}) {
	l.logger.Infof(strings.Repeat("%+v ", len(m)), m...)
}

func (l *opLogger) Error(ctx context.Context, m ...interface{}) {
	l.logger.Errorf(strings.Repeat("%+v ", len(m)), m...)
}

func (l *opLogger) Debug(ctx context.Context, m ...interface{}) {
	l.logger.Debugf(strings.Repeat("%+v ", len(m)), m...)
}
