package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityTypeServer) List(ctx context.Context, req *activitytypecomm.ListActivityTypesRequest) (*activitytypecomm.ListActivityTypesReply, error) {
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

	activityTypes := []*activitytypecomm.ActivityType{}
	for _, v := range acts {
		activityTypes = append(activityTypes, &activitytypecomm.ActivityType{
			Code: v.Code,
			Name: v.Name,
			Id:   v.ID,
		})
	}

	return &activitytypecomm.ListActivityTypesReply{
		ActivityTypes: activityTypes,
	}, nil
}
