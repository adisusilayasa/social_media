package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"social_media/config"
	"social_media/pkg/zap"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const (
	MaxHeaderBytes = 1 << 20
)

type Server struct {
	cfg    *config.Config
	logger zap.Logger
	db     *gorm.DB
	echo   *echo.Echo
}

func NewServer(cfg *config.Config, logger zap.Logger, db *gorm.DB) *Server {
	return &Server{cfg: cfg, logger: logger, db: db, echo: echo.New()}
}

func (s *Server) Run() error {

	s.echo.HideBanner = true
	s.echo.HidePort = true

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	server := &http.Server{
		Handler:        s.echo,
		Addr:           "localhost:8080",
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: 8080")
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatalf("Error starting server: %v", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
