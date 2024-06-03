package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debub  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debub:  log.New(writer, "[DEBUG] ", logger.Flags()),
		info:   log.New(writer, "[INFO] ", logger.Flags()),
		warn:   log.New(writer, "[WARN] ", logger.Flags()),
		error:  log.New(writer, "[ERROR] ", logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debub.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debub.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.error.Printf(format, v...)
}
