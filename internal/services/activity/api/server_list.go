package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) List(ctx context.Context, req *activitycomm.ListActivitiesRequest) (*activitycomm.ListActivitiesReply, error) {
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
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	// collecting the needed ActivityType details
	var actTypes map[string]*activitytypecomm.ActivityType
	if len(acts) > 0 {
		types := []string{}
		for _, act := range acts {
			types = append(types, act.ActivityTypeID)
		}

		activityTypesReply, err := s.actTypeCli.List(ctx, &activitytypecomm.ListActivityTypesRequest{FilterIds: types})
		if err != nil {
			log.Printf("error retrieving the filtered list of activity types: %v", err)
		} else {
			actTypes = make(map[string]*activitytypecomm.ActivityType, len(activityTypesReply.ActivityTypes))
			for _, at := range activityTypesReply.ActivityTypes {
				actTypes[at.Id] = at
			}
		}
	}

	// building response with optional data
	activities := make([]*activitycomm.Activity, len(acts))
	for idx, act := range acts {
		// if info about the ActivityType was not retrieved, add the info you have
		actType, ok := actTypes[act.ActivityTypeID]
		if !ok {
			actType = &activitytypecomm.ActivityType{Id: act.ActivityTypeID}
		}

		from, _ := ptypes.TimestampProto(act.Period.From)
		to, _ := ptypes.TimestampProto(act.Period.To)
		activities[idx] = &activitycomm.Activity{
			Period: &activitycomm.Interval{
				From: from,
				To:   to,
			},
			Id:      act.ID,
			ActType: actType,
		}
	}

	return &activitycomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}
