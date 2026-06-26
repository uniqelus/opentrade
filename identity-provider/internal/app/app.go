package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	sdkgrpc "github.com/uniqelus/opentrade/sdk/go/grpc"
	sdklog "github.com/uniqelus/opentrade/sdk/go/log/slog"
	"golang.org/x/sync/errgroup"

	userrepo "github.com/uniqelus/opentrade/identity-provider/internal/repositories/user"
	usersrv "github.com/uniqelus/opentrade/identity-provider/internal/services/user"
	userapi "github.com/uniqelus/opentrade/identity-provider/internal/transport/grpc/user"
)

type App struct {
	log *slog.Logger

	pgxPool    *pgxpool.Pool
	grpcServer *sdkgrpc.Server
}

func NewApp(cfg *Config) (*App, error) {
	log, err := sdklog.NewLog(cfg.Log.Options()...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize app: %w", err)
	}

	log = log.With(slog.String("application", "identity-provider"))
	log.Info("application initialized")

	log.Debug("parsing database configuration")
	poolConfig, err := pgxpool.ParseConfig(cfg.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database dsn: %w", err)
	}

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig) // TODO: fix context.Background()
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	if err := pgxPool.Ping(context.Background()); err != nil { // TODO: fix context.Background()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Info("successfully connected to database")

	userRepo := userrepo.NewRepository(pgxPool)
	userService := usersrv.NewService(userRepo)

	log.Debug("initialize grpc server",
		slog.String("host", cfg.GRPCServer.Host),
		slog.String("port", cfg.GRPCServer.Port),
	)

	grpcServerOptions := append(cfg.GRPCServer.Options(),
		sdkgrpc.WithServiceRegistrations(userapi.NewServiceRegistration(userService)),
	)
	grpcServer := sdkgrpc.NewServer(grpcServerOptions...)

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
