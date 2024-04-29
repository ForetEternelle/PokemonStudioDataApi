package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/studio"
)

func main() {
	config := ParseApiConfig()
	store := studio.NewStore()

	if err := studio.Import(config.DataFolder, store); err != nil {
		panic("Failed to importe data folder")
	}
	psapiRouter := psapi.MakeDefaultRouter(store)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", config.Cors))
	r.Mount("/", psapiRouter)

	addr := fmt.Sprintf(":%d", config.Port)
	slog.Info("Server listening", "addr", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	server.ListenAndServe()
}
