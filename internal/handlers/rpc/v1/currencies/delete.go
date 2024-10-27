package currencies

import (
	"context"
	pb "exporterbackend/proto/gen/app/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Delete(ctx context.Context, req *pb.DeleteCurrencyRequest) (*pb.CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
