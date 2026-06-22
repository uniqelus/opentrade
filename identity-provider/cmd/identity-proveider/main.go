package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/uniqelus/opentrade/identity-provider/internal/app"
	sdkconfig "github.com/uniqelus/opentrade/sdk/go/config"
	sdkerrors "github.com/uniqelus/opentrade/sdk/go/errors"
)

func main() {

	var configPath string

	flag.StringVar(&configPath, "config", "", "path to configuration file")
	flag.Parse()

	cfg := sdkerrors.Must(sdkconfig.Read[app.Config](configPath))
	app := sdkerrors.Must(app.NewApp(cfg))

	runCtx, runCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer runCancel()

	if err := app.Run(runCtx); err != nil {
		panic(err)
	}

	<-runCtx.Done()
	runCancel()

	stopCtx, stopCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopCancel()

	if err := app.Stop(stopCtx); err != nil {
		panic(err)
	}
}
