package currencies

import (
	"context"
	pb "exporterbackend/proto/gen/app/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Create(ctx context.Context, req *pb.CreateCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
