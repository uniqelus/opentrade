package app

import (
	"context"
	"fmt"
	"log/slog"

	sdkgrpc "github.com/uniqelus/opentrade/sdk/go/grpc"
	sdklog "github.com/uniqelus/opentrade/sdk/go/log/slog"
	"golang.org/x/sync/errgroup"
)

type App struct {
	log *slog.Logger

	grpcServer *sdkgrpc.Server
}

func NewApp(cfg *Config) (*App, error) {
	log, err := sdklog.NewLog(cfg.Log.Options()...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize app: %w", err)
	}

	log = log.With(slog.String("application", "identity-provider"))
	log.Info("application initialized")

	log.Debug("initialize grpc server",
		slog.String("host", cfg.GRPCServer.Host),
		slog.String("port", cfg.GRPCServer.Port),
	)
	grpcServer := sdkgrpc.NewServer(cfg.GRPCServer.Options()...)

	return &App{
		log:        log,
		grpcServer: grpcServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("starting application")

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Debug("running grpc server")
		return a.grpcServer.Start(gCtx)
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
		a.log.Debug("stopping grpc server")
		return a.grpcServer.Stop(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("critical error occurred while stopping the application", sdklog.Error(err))
		return fmt.Errorf("critical error occurred while stopping the application: %w", err)
	}

	return nil
}
