package timber

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// NewGoFileLogger returns new logger
func NewGoFileLogger(name string) Logger {
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

	files := logFile
	//files := io.MultiWriter(logFile, os.Stdout)

	logger := log.New(files, "", log.LstdFlags|log.LUTC)

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
