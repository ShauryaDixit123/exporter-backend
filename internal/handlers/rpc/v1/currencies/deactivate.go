package currencies

import (
	"context"
	pb "exporterbackend/proto/gen/app/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Deactivate(ctx context.Context, req *pb.DeactiviateCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
