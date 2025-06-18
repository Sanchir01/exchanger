package app

import (
	"github.com/Sanchir01/exchanger/internal/feature/exchanger"
	"google.golang.org/grpc"
	"log/slog"
)

type Handler struct {
	ExchangerHandler *exchanger.Handler
}

func NewHandlers(srv *Services, server *grpc.Server, l *slog.Logger) *Handler {
	return &Handler{
		ExchangerHandler: exchanger.NewHandler(srv.ExchangerService, server, l),
	}
}
