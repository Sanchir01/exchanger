package main

import (
	"context"
	"fmt"
	"github.com/Sanchir01/exchanger/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	env, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(env.Cfg)
	env.Lg.Info("Starting exchanger")

	defer stop()
	go func() {
		env.GRPCSrv.MustRun()
	}()
	<-ctx.Done()
	env.Lg.Warn("Stopping exchanger")
	env.GRPCSrv.Stop()
}
