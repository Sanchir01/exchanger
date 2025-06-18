package exchanger

import (
	"context"
	"fmt"
	walletsv1 "github.com/Sanchir01/wallets-proto/gen/go/wallets"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

type Handler struct {
	service Service
	walletsv1.UnimplementedExchangeServiceServer
	log *slog.Logger
}

type Service interface {
	AllCurrency(ctx context.Context) ([]ExchangeDB, error)
	GetExchangeRate(ctx context.Context, fromCurrency, toCurrency string) (*ExchangeDB, error)
}

func NewHandler(service Service, server *grpc.Server, l *slog.Logger) *Handler {
	handler := &Handler{service: service, log: l}
	walletsv1.RegisterExchangeServiceServer(server, handler)
	return handler
}
func (h *Handler) GetExchangeRates(ctx context.Context, empty *emptypb.Empty) (*walletsv1.ExchangeRatesResponse, error) {
	const op = "handlers.getExchangeRateForCurrency"
	log := h.log.With(
		slog.String("op", op),
	)
	data, err := h.service.AllCurrency(ctx)
	if err != nil {
		log.Error("GetExchangeRates", "error", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	rates := make(map[string]float32)

	for _, item := range data {
		key := fmt.Sprintf("%s/%s", item.FromCurrency, item.ToCurrency)
		rates[key] = float32(item.Rate)
	}
	log.Info("success handler get all exchange rates")
	return &walletsv1.ExchangeRatesResponse{
		Rates: rates,
	}, nil
}

func (h *Handler) GetExchangeRateForCurrency(ctx context.Context, request *walletsv1.CurrencyRequest) (*walletsv1.ExchangeRateResponse, error) {
	const op = "handlers.getExchangeRateForCurrency"
	log := h.log.With(
		slog.String("op", op),
	)
	req := GetExchangeRate{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
	if err := validator.New().Struct(req); err != nil {
		log.Error("failed to validate request", "error", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	data, err := h.service.GetExchangeRate(ctx, request.FromCurrency, request.ToCurrency)
	if err != nil {
		log.Error("GetExchangeRateForCurrency", "error", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info("success handler get data")

	return &walletsv1.ExchangeRateResponse{
		Rate:         float32(data.Rate),
		ToCurrency:   data.ToCurrency,
		FromCurrency: data.FromCurrency,
	}, nil
}
