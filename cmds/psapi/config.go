package main

import (
	"flag"
	"log/slog"
)

const (
	AppConfigDefaultName = "psapi"

	KeyAppLogLevel     = "log-level"
	DefaultAppLogLevel = "INFO"

	KeyApiCors     = "cors"
	DefaultApiCors = "*"

	KeyApiPort     = "port"
	DefaultApiPort = 8000

	KeyImportDataFolderPath     = "data"
	DefaultImportDataFolderPath = "data"
)

type Config struct {
	LogLevel   slog.Level
	DataFolder string
	Port       int
	Cors       string
}

func ParseApiConfig() *Config {
	logLevelStr := flag.String(KeyAppLogLevel, DefaultAppLogLevel, "The log level")
	dataFolder := flag.String(KeyImportDataFolderPath, DefaultImportDataFolderPath, "Data folder")
	port := flag.Int(KeyApiPort, DefaultApiPort, "port to serve server on")
	cors := flag.String(KeyApiCors, DefaultApiCors, "cors header")
	flag.Parse()

	logLevel := ParseLogLevel(*logLevelStr)

	return &Config{
		LogLevel:   logLevel,
		DataFolder: *dataFolder,
		Port:       *port,
		Cors:       *cors,
	}
}

func ParseLogLevel(levelStr string) slog.Level {
	var level slog.Level
	switch levelStr {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	return level
}
