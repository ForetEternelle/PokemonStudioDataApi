package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	customMiddleware "github.com/ForetEternelle/PokemonStudioDataApi/pkg/middleware"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/pkmnapi"
	"github.com/ForetEternelle/PokemonStudioDataApi/pkg/pkmn/studio"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	AppConfigDefaultName = "psapi"

	KeyAppLogLevel     = "log-level"
	DefaultAppLogLevel = "INFO"

	KeyApiCors     = "cors"
	DefaultApiCors = "*"

	KeyApiPort     = "port"
	DefaultApiPort = 8080

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
	logLevelStr := flag.String(KeyAppLogLevel, envOr("LOG_LEVEL", DefaultAppLogLevel), "The log level")
	dataFolder := flag.String(KeyImportDataFolderPath, envOr("DATA", DefaultImportDataFolderPath), "Data folder")
	port := flag.Int(KeyApiPort, envOrInt("PORT", DefaultApiPort), "port to serve server on")
	cors := flag.String(KeyApiCors, envOr("CORS", DefaultApiCors), "cors header")
	flag.Parse()

	logLevel := ParseLogLevel(*logLevelStr)

	return &Config{
		LogLevel:   logLevel,
		DataFolder: *dataFolder,
		Port:       *port,
		Cors:       *cors,
	}
}

func envOr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func envOrInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
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

	studioApiRouter, err := pkmnapi.GetRouter(
		pkmnapi.WithStore(store),
	)
	if err != nil {
		panic("Failed to create api router. Error: " + err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", config.Cors))

	startDate := time.Now().UTC()
	r.Use(middleware.SetHeader("Last-Modified", startDate.Format(http.TimeFormat)))
	r.Use(customMiddleware.Cache(startDate))
	r.Mount("/api", studioApiRouter)

	addr := fmt.Sprintf(":%d", config.Port)
	slog.Info("Server listening", "addr", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	server.ListenAndServe()
}
