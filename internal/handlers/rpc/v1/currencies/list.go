package currencies

import (
	"context"

	pb "exporterbackend/proto/gen/app/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) List(ctx context.Context, req *pb.ListCurrenciesRequest) (*pb.CurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
