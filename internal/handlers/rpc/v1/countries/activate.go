package countries

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "exporterbackend/proto/gen/app/v1"
)

func (h *Handler) Activate(ctx context.Context, req *pb.ActiviateCountryRequest) (*pb.CountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
