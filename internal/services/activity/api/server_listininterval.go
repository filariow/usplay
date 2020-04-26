package api

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) ListInInterval(ctx context.Context, req *activitycomm.ListInIntervalActivitiesRequest) (*activitycomm.ListActivitiesReply, error) {
	periodComm := req.GetPeriod()
	if periodComm == nil {
		return nil, status.Errorf(codes.InvalidArgument, "error retrieving the list of activities: interval not provided")
	}

	period, err := storage.NewIntervalProto(periodComm.From, periodComm.To)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error retrieving the list of activities: interval not valid: %v", err)
	}
	acts, err := s.repo.ListInInterval(ctx, *period)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving the list of activities: %v", err)
	}

	// collecting the needed ActivityType details
	var actTypes map[string]*activitytypecomm.ActivityType
	if len(acts) > 0 {
		_types := map[string]*struct{}{}
		for _, act := range acts {
			_types[act.ActivityTypeID] = nil
		}

		counter, types := 0, make([]string, len(_types))
		for k := range _types {
			types[counter] = k
			counter++
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
			Id:      act.ID,
			ActType: actType,
			Period: &activitycomm.Interval{
				From: from,
				To:   to,
			},
		}
	}

	return &activitycomm.ListActivitiesReply{
		Activities: activities,
	}, nil
}
