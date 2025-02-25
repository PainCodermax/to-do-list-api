package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/PainCodermax/to-do-list-api/pkg/setting"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *lumberjack.Logger

func SetupLogger(cfg *setting.Configuration) {
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(cfg.LogSavePath, cfg.LogFileName),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     3,
		Compress:   cfg.Compress,
	}
	lvl, err := log.ParseLevel(cfg.Level)
	if err != nil {
		log.SetLevel(lvl)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	mw := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(mw)

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Close() {
	logger.Close()
}

func WithTrace(ctx *gin.Context) *log.Entry {
	fields := log.Fields{}
	if len(ctx.GetString("X-Trace-ID")) > 0 {
		fields["trace_id"] = ctx.GetString("X-Trace-ID")
	}
	if len(ctx.GetString("X-Span-ID")) > 0 {
		fields["span_id"] = ctx.GetString("X-Span-ID")
	}
	return log.WithFields(fields)
}
