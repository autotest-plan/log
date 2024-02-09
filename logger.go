package log

import (
	"go.uber.org/zap"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func (l *Logger) Error(args ...interface{}) { l.sugar.Error(args) }

func (l *Logger) Errorf(template string, args ...interface{}) { l.sugar.Errorf(template, args) }

func (l *Logger) Errorln(args ...interface{}) { l.sugar.Errorln(args) }

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) { l.sugar.Errorw(msg, keysAndValues) }

func (l *Logger) Info(args ...interface{}) { l.sugar.Info(args) }

func (l *Logger) Infof(template string, args ...interface{}) { l.sugar.Infof(template, args) }

func (l *Logger) Infoln(args ...interface{}) { l.sugar.Infoln(args) }

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) { l.sugar.Infow(msg, keysAndValues) }

func (l *Logger) Debug(args ...interface{}) { l.sugar.Debug(args) }

func (l *Logger) Debugf(template string, args ...interface{}) { l.sugar.Debugf(template, args) }

func (l *Logger) Debugln(args ...interface{}) { l.sugar.Debugln(args) }

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) { l.sugar.Debugw(msg, keysAndValues) }

func (l *Logger) Warn(args ...interface{}) { l.sugar.Warn(args) }

func (l *Logger) Warnf(template string, args ...interface{}) { l.sugar.Warnf(template, args) }

func (l *Logger) Warnln(args ...interface{}) { l.sugar.Warnln(args) }

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) { l.sugar.Warnw(msg, keysAndValues) }

func (l *Logger) Fatal(args ...interface{}) { l.sugar.Fatal(args) }

func (l *Logger) Fatalf(template string, args ...interface{}) { l.sugar.Fatalf(template, args) }

func (l *Logger) Fatalln(args ...interface{}) { l.sugar.Fatalln(args) }

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) { l.sugar.Fatalw(msg, keysAndValues) }

func NewDevelopmentLogger(path []string) (*Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = append(config.OutputPaths, path...)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	return &Logger{sugar: sugar}, nil
}

func NewProductLogger(path []string) (*Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, path...)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()
	return &Logger{sugar: sugar}, nil
}
