package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log logger
)

type bookstorelogger interface {
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// ElasticSearch Logger Implementation
// Logger specifies the interface for all log operations.
// type Logger interface {
// 	Printf(format string, v ...interface{})
// }

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		l.Info(format)
	} else {
		l.Info(fmt.Sprintf(format, v...))
	}
}

func GetLogger() bookstorelogger {
	return log
}

func (l logger) Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func (l logger) Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
