package app

import (
	"context"
	"github.com/Sanchir01/exchanger/internal/config"
	"github.com/Sanchir01/exchanger/pkg/db"
	"github.com/Sanchir01/exchanger/pkg/logger"
	grpcapp "github.com/Sanchir01/exchanger/pkg/server/grpc"
	"log/slog"
)

type App struct {
	Cfg     *config.Config
	Lg      *slog.Logger
	DB      *db.Database
	GRPCSrv *grpcapp.App
}

func NewApp(ctx context.Context) (*App, error) {
	cfg := config.InitConfig()
	l := logger.SetupLogger(cfg.Env)
	database, err := db.NewDataBases(cfg, ctx)
	if err != nil {
		l.Error("db error connect", err.Error())
		return nil, err
	}
	repo := NewRepository(database)
	srv := NewServices(repo, l)
	gRPCSrv := grpcapp.GetGrpcServer()
	_ = NewHandlers(srv, gRPCSrv, l)
	gRPCServer := grpcapp.New(l, ":"+cfg.GRPC.Port, gRPCSrv)
	return &App{
		Cfg:     cfg,
		Lg:      l,
		DB:      database,
		GRPCSrv: gRPCServer,
	}, nil
}
