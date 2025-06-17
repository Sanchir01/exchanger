package exchanger

import (
	"context"
	walletsv1 "github.com/Sanchir01/wallets-proto/gen/go/wallets"
	"google.golang.org/grpc"
)

type Handler struct {
	service Service
	walletsv1.UnimplementedExchangeServiceServer
}

type Service interface {
}

func NewHandler(service Service, server *grpc.Server) *Handler {
	handler := &Handler{service: service}
	walletsv1.RegisterExchangeServiceServer(server, handler)
	return handler
}
func (h Handler) GetExchangeRates(ctx context.Context, empty *walletsv1.Empty) (*walletsv1.ExchangeRatesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetExchangeRateForCurrency(ctx context.Context, request *walletsv1.CurrencyRequest) (*walletsv1.ExchangeRateResponse, error) {
	//TODO implement me
	panic("implement me")
}
