package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/bookmastergrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) List(ctx context.Context, req *bookmastergrpc.ListActivityTypesRequest) (*bookmastergrpc.ListActivityTypesReply, error) {
	ids := make([]uuid.UUID, len(req.FilterIds))
	for idx, i := range req.FilterIds {
		id, err := uuid.Parse(i)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "ID is not valid: %v", err)
		}
		ids[idx] = id
	}

	acts, err := s.repo.List(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of ActivityTypes: %v", err)
	}

	activityTypes := []*bookmastergrpc.ActivityType{}
	for _, v := range acts {
		activityTypes = append(activityTypes, &bookmastergrpc.ActivityType{
			Code: v.Code,
			Name: v.Name,
			Id:   v.ID,
		})
	}

	return &bookmastergrpc.ListActivityTypesReply{
		ActivityTypes: activityTypes,
	}, nil
}
