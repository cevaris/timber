package timber

import (
	"context"
	"strings"

	"google.golang.org/appengine/log"
)

type gaeLoggger struct {
}

// NewAppEngineLogger is an google app engine backed logger
func NewAppEngineLogger() Logger {
	return &gaeLoggger{}
}

func (l *gaeLoggger) Info(ctx context.Context, m ...interface{}) {
	if Level == Info || Level == Error {
		log.Infof(ctx, strings.Repeat("%v ", len(m)), m...)
	}
}

func (l *gaeLoggger) Error(ctx context.Context, m ...interface{}) {
	// always log errors
	log.Errorf(ctx, strings.Repeat("%v ", len(m)), m...)
}

func (l *gaeLoggger) Debug(ctx context.Context, m ...interface{}) {
	if Level == Debug {
		log.Debugf(ctx, strings.Repeat("%v ", len(m)), m...)
	}
}
