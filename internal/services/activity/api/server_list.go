package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) List(ctx context.Context, req *activitycomm.ListActivitiesRequest) (*activitycomm.ListActivitiesReply, error) {
	acts, err := s.repo.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	// collecting the needed ActivityType details
	var actTypes map[string]*activitytypecomm.ActivityType
	if len(acts) > 0 {
		types := []string{}
		for _, act := range acts {
			types = append(types, act.ActivityTypeID.String())
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
		actType, ok := actTypes[act.ActivityTypeID.String()]
		if !ok {
			actType = &activitytypecomm.ActivityType{Id: act.ActivityTypeID.String()}
		}

		activities[idx] = &activitycomm.Activity{
			Code:        act.Code,
			Description: act.Description,
			Name:        act.Name,
			Id:          act.ID.String(),
			ActType:     actType,
		}
	}

	return &activitycomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}
