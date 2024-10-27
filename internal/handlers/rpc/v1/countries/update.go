package countries

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "exporterbackend/proto/gen/app/v1"
)

func (h *Handler) Update(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.CountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
