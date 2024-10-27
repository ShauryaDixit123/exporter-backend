package currencies

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "exporterbackend/proto/gen/app/v1"
)

func (h *Handler) Activate(ctx context.Context, req *pb.ActiviateCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
