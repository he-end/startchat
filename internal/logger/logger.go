package logger

import (
	"log"
	"os"

	"github.com/hend41234/startchat/internal/internalutils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Environment     string // "development" | "production"
	LogToConsole    bool
	LogToFile       bool
	LogToRemote     bool
	EnableRolling   bool
	LogFilePath     string
	MinimumLogLevel string // "debug", "info", "warn", "error"
}

var Log *zap.Logger

func Init(cfg Config) {
	var cores []zapcore.Core

	// Set encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.CallerKey = "caller"

	// Set log level
	level := zapcore.InfoLevel
	_ = level.UnmarshalText([]byte(cfg.MinimumLogLevel))

	// ========== File Logging ==========
	if cfg.LogToFile {
		err := internalutils.EnsureDir(cfg.LogFilePath)
		if err != nil {
			log.Fatalf("failed to create log directory: %v", err)
		}

		var writer zapcore.WriteSyncer

		if cfg.EnableRolling {
			writer = zapcore.AddSync(&lumberjack.Logger{
				Filename:   cfg.LogFilePath,
				MaxSize:    10, // MB
				MaxBackups: 5,
				MaxAge:     7,    // days
				Compress:   true, // gzip
			})
		} else {
			file, err := os.OpenFile(cfg.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("failed to open log file: %v", err)
			}
			writer = zapcore.AddSync(file)
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			writer,
			level,
		)
		cores = append(cores, core)
	}

	// ========== Console Logging ==========
	if cfg.LogToConsole {
		var consoleEncoder zapcore.Encoder
		if cfg.Environment == "development" {
			consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		} else {
			consoleEncoder = zapcore.NewConsoleEncoder(encoderCfg)
		}

		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	// ========== Remote Logging (future use) ==========
	if cfg.LogToRemote {
		// TODO: Implement remote logger writer (e.g., HTTP, Kafka, etc.)
		// remoteWriter := NewRemoteWriter()
		// remoteCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(remoteWriter), level)
		// cores = append(cores, remoteCore)
	}

	if len(cores) == 0 {
		log.Fatal("No logging output configured. Enable at least one of: LogToConsole, LogToFile, LogToRemote.")
	}

	core := zapcore.NewTee(cores...)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}
