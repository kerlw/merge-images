package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func createLogger() *zap.Logger {
	var alevel zap.AtomicLevel
	alevel.UnmarshalText([]byte("debug"))
	fmt.Printf("log.level=%v\n", alevel)
	cfg := zap.Config{
		Level:            alevel,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			TimeKey:     "time",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02T15:04:05"))
			},
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Printf("failed to create logger: %s", err.Error())
		os.Exit(-1)
	}

	return logger
}
