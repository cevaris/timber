package timber

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/logging"
)

type gceLoggger struct {
	logger *logging.Logger
}

// NewComputeEngineLogger is an google app engine backed logger
func NewComputeEngineLogger(ctx context.Context, projectID string, name string) Logger {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return &gceLoggger{logger: client.Logger(name)}
}

func (l *gceLoggger) Info(ctx context.Context, m ...interface{}) {
	if Level == Info || Level == Error {
		entry := logging.Entry{Payload: fmt.Sprintf(strings.Repeat("%v ", len(m)), m...)}
		l.logger.Log(entry)
	}
}

func (l *gceLoggger) Error(ctx context.Context, m ...interface{}) {
	// always log errors
	entry := logging.Entry{Payload: fmt.Sprintf(strings.Repeat("%v ", len(m)), m...)}
	l.logger.Log(entry)
}

func (l *gceLoggger) Debug(ctx context.Context, m ...interface{}) {
	if Level == Debug {
		entry := logging.Entry{Payload: fmt.Sprintf(strings.Repeat("%v ", len(m)), m...)}
		l.logger.Log(entry)
	}
}
