package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(level zapcore.Level, debug bool, outputPath string) (*zap.SugaredLogger, *zap.AtomicLevel, error) {
	if len(outputPath) == 0 {
		outputPath = "stdout"
	}
	lev := zap.NewAtomicLevelAt(level)
	var c zap.Config
	if debug {
		c = zap.Config{
			Development:       false,     // production
			DisableCaller:     false,     // always print function name and line number
			DisableStacktrace: false,     // do not print stacktrace
			Sampling:          nil,       // print all log
			Encoding:          "console", // not json
			EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
			OutputPaths:       []string{outputPath},
			ErrorOutputPaths:  []string{outputPath},
		}
		c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		c = zap.NewProductionConfig()
		c.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		c.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}
	// logger.level is a pointer, so it can be changed dynamically.
	// if logger.Named() is called, although it creates a new logger, but they still share a same level pointer
	// so change level will cause all loggers which have the same root logger changed their level
	c.Level = lev
	l, err := c.Build()
	if err != nil {
		return nil, nil, fmt.Errorf("build logger failed: %w", err)
	}
	return l.Sugar(), &lev, nil
}
