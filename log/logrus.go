package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerCtxKey string

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func InfoWithFields(fields map[string]interface{}, args ...interface{}) {
	logrus.WithFields(fields).Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func InfofWithFieldsCtx(ctx context.Context, fields map[string]interface{}, format string, args ...interface{}) {
	logrus.WithFields(fields).Infof(format, args...)
}

func InfofWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	logrus.WithFields(fields).Infof(format, args...)
}

func Warn(err error, args ...interface{}) {
	logrus.Warn(args...)
}

func WarnWithFields(err error, fields map[string]interface{}, args ...interface{}) {
	logrus.WithFields(fields).Warn(args...)
}

func WarnfWithFields(err error, fields map[string]interface{}, format string, args ...interface{}) {
	logrus.WithFields(fields).Warnf(format, args...)
}

func Error(err error, args ...interface{}) {
	logrus.WithField("status", "error").Error(args...)
}

func ErrorWithFields(fields map[string]interface{}, err error, args ...interface{}) {
	logrus.WithField("status", "error").WithFields(fields).Error(args...)
}

func Errorf(err error, format string, args ...interface{}) {
	logrus.WithField("status", "error").Errorf(format, args...)
}

func ErrorfWithFields(fields map[string]interface{}, err error, format string, args ...interface{}) {
	logrus.WithField("status", "error").WithFields(fields).Errorf(format, args...)
}

func Warnf(err error, format string, args ...interface{}) {
	logrus.WithField("status", "warn").Warnf(format, args...)
}

func Fatal(err error, args ...interface{}) {
	logrus.WithField("status", "error").Fatal(args...)
}

func FatalWithFields(fields map[string]interface{}, err error, args ...interface{}) {
	logrus.WithField("status", "error").WithFields(fields).Fatal(args...)
}

func Fatalf(format string, err error, args ...interface{}) {
	logrus.WithField("status", "error").Fatalf(format, args...)
}

func FatalfWithFields(fields map[string]interface{}, err error, format string, args ...interface{}) {
	logrus.WithField("status", "error").WithFields(fields).Fatalf(format, args...)
}
