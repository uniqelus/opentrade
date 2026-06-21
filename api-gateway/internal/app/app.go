package app

import (
	"context"
	"fmt"
	"log/slog"

	sdkhttp "github.com/uniqelus/opentrade/sdk/go/http"
	sdklog "github.com/uniqelus/opentrade/sdk/go/log/slog"
	"golang.org/x/sync/errgroup"

	httptr "github.com/uniqelus/opentrade/api-gateway/internal/transport/http"
)

type App struct {
	log *slog.Logger

	httpServer *sdkhttp.Server
}

func NewApp(cfg *Config) (*App, error) {
	log, err := sdklog.NewLog(cfg.Log.Options()...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize app: %w", err)
	}

	log = log.With(slog.String("application", "api-gateway"))
	log.Info("application initialized")

	log.Debug("initialize http server",
		slog.String("host", cfg.HTTPServer.Host),
		slog.String("port", cfg.HTTPServer.Port),
		slog.Duration("read_duration", cfg.HTTPServer.ReadTimeout),
		slog.Duration("write_duration", cfg.HTTPServer.WriteTimeout),
		slog.Duration("idle_duration", cfg.HTTPServer.IdleTimeout),
	)
	httpServerOptions := append(cfg.HTTPServer.Options(), sdkhttp.WithServerHandler(httptr.NewRouter()))
	httpServer := sdkhttp.NewServer(httpServerOptions...)

	return &App{
		log:        log,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("starting application")

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Debug("running http server")
		return a.httpServer.Run(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("critical error occurred while the application was running", sdklog.Error(err))
		return fmt.Errorf("critical error occurred while the application was running: %w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.log.Info("stopping application")

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Debug("stopping http server")
		return a.httpServer.Stop(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("critical error occurred while stopping the application", sdklog.Error(err))
		return fmt.Errorf("critical error occurred while stopping the application: %w", err)
	}

	return nil
}
