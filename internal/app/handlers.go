package app

import (
	"github.com/Sanchir01/exchanger/internal/feature/exchanger"
	"google.golang.org/grpc"
)

type Handler struct {
	ExchangerHandler *exchanger.Handler
}

func NewHandlers(srv *Services, server *grpc.Server) *Handler {
	return &Handler{
		ExchangerHandler: exchanger.NewHandler(srv, server),
	}
}
