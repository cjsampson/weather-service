package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerState struct {
	api *http.Server
	log *slog.Logger

	SignalCh chan os.Signal
	ErrorCh  chan error
}

// NewServerState constructor - TODO pass in configuration
func NewServerState(mux *http.ServeMux, log *slog.Logger) *ServerState {
	return &ServerState{
		api: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		SignalCh: make(chan os.Signal, 1),
		ErrorCh:  make(chan error, 1),
		log:      log,
	}
}

// Start initializes the server and listens for shutdown signals.
func (s *ServerState) Start() error {
	s.log.Info("startup", slog.String("status", "initializing"))

	signal.Notify(s.SignalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		s.log.Info("Application started", slog.String("host", s.api.Addr))
		s.ErrorCh <- s.api.ListenAndServe()
	}()

	return s.waitForSignal()
}

// waitForSignal handles shutdown or error signals.
func (s *ServerState) waitForSignal() error {
	select {
	case err := <-s.ErrorCh:
		return fmt.Errorf("server error: %w", err)

	case sig := <-s.SignalCh:
		return s.gracefulShutdown(sig)
	}
}

// gracefulShutdown attempts to gracefully shut down the server.
func (s *ServerState) gracefulShutdown(sig os.Signal) error {
	s.log.Info("shutdown started", slog.Any("signal", sig))
	defer s.log.Info("shutdown finished", slog.Any("signal", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := s.api.Shutdown(ctx); err != nil {
		_ = s.api.Close()
		return fmt.Errorf("could not stop server gracefully: %w", err)
	}
	return nil
}
