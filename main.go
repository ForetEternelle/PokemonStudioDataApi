package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/studio/studioapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	AppConfigDefaultName = "psapi"

	KeyAppLogLevel     = "log-level"
	DefaultAppLogLevel = "DEBUG"

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

func main() {
	config := ParseApiConfig()
	slog.SetLogLoggerLevel(config.LogLevel)
	store, err := studio.Load(config.DataFolder)

	if err != nil {
		panic("Failed to import data folder. Error: " + err.Error())
	}

	studioApiRouter := studioapi.MakeDefaultRouter(store)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", config.Cors))
	r.Mount("/api", studioApiRouter)

	addr := fmt.Sprintf(":%d", config.Port)
	slog.Info("Server listening", "addr", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	server.ListenAndServe()
}
