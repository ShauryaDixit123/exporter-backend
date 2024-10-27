package countries

import (
	"context"
	pb "exporterbackend/proto/gen/app/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Create(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
