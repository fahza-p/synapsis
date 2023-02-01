package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
	"go.opentelemetry.io/otel/trace"
)

func SetLevel(level string) {
	switch level {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func GetLevel() string {
	return strings.ToUpper(logrus.GetLevel().String())
}

func GetLogger(ctx context.Context, pkg, fnName string) *logrus.Entry {

	_, file, _, _ := runtime.Caller(1)
	file = file[strings.LastIndex(file, "/")+1:]

	fields := logrus.Fields{
		"function": fnName,
		"package":  pkg,
		"source":   file,
	}
	span := trace.SpanFromContext(ctx)
	if span != nil {
		if span.SpanContext().HasSpanID() {
			fields["span_id"] = span.SpanContext().SpanID().String()
		}
		if span.SpanContext().HasTraceID() {
			fields["trace_id"] = span.SpanContext()
		}
	}
	return WithContext(ctx).WithFields(fields)
}

func WithContext(ctx context.Context) *logrus.Entry {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	logrus.SetReportCaller(true)
	return logrus.WithContext(ctx).WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func Configure(format, level string) {
	switch strings.ToLower(format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	lvl := strings.ToLower(level)
	SetLevel(lvl)
	levels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}

	if lvl == "debug" {
		levels = append(levels, logrus.DebugLevel)
	}

	logrus.AddHook(otellogrus.NewHook(otellogrus.WithLevels(levels...)))
}
