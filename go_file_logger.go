package timber

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// NewGoFileLogger returns new logger
// if stdout==true, multi-write logs to stdout and file
func NewGoFileLogger(name string, stdout bool) Logger {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("could locate $USER home directory: " + err.Error())
	}

	logDirPath := filepath.Join(homeDir, "var", "log", "status")
	if err = os.MkdirAll(logDirPath, os.ModePerm); err != nil {
		panic("could create log directory: " + err.Error())
	}

	logFilePath := filepath.Join(logDirPath, name+".log")
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("could create log file for " + name + ": " + err.Error())
	}

	fmt.Println("created log file", logFilePath)

	var sink io.Writer
	if stdout {
		sink = io.MultiWriter(logFile, os.Stdout)
	} else {
		sink = logFile
	}

	logger := log.New(sink, "", log.LstdFlags|log.LUTC)

	return &goFileLogger{logger: logger}
}

type goFileLogger struct {
	logger *log.Logger
}

func (l *goFileLogger) Info(ctx context.Context, m ...interface{}) {
	if Level == Error || Level == Info {
		l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
	}
}

func (l *goFileLogger) Error(ctx context.Context, m ...interface{}) {
	l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
}

func (l *goFileLogger) Debug(ctx context.Context, m ...interface{}) {
	if Level == Debug {
		l.logger.Printf(strings.Repeat("%+v ", len(m))+"\n", m...)
	}
}
