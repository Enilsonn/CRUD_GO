package main

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Enilsonn/CRUD_GO/api"
	"github.com/Enilsonn/CRUD_GO/api/handlers"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code",
			"error", err,
		)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	bd := handlers.BD{
		Users: make(map[int]handlers.User),
		Len:   0,
	}

	mux := api.NewMultiplier(&bd)

	s := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
