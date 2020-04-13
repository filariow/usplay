package api

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *activityServer) Read(ctx context.Context, req *activitycomm.ReadActivityRequest) (*activitycomm.ReadActivityReply, error) {
	id := req.GetId()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id provided (%s): %v", id, err)
	}

	return s.read(ctx, uid)
}

func (s *activityServer) read(ctx context.Context, uid uuid.UUID) (*activitycomm.ReadActivityReply, error) {
	bgCtx, bgCancFunc := context.WithTimeout(ctx, s.waitTime)
	cat := make(chan activitytypecomm.ActivityType)
	ca := make(chan activitycomm.Activity)

	go func() {
		if act, err := s.getActivityType(bgCtx, uid); err != nil {
			bgCancFunc()
		} else {
			cat <- *act
		}
	}()

	go func() {
		if at, err := s.readFromRepo(bgCtx, uid); err != nil {
			bgCancFunc()
		} else {
			ca <- *at
		}
	}()

	var act activitycomm.Activity
	select {
	case act = <-ca:
		select {
		case acttype := <-cat:
			act.ActType = &acttype
			return &activitycomm.ReadActivityReply{
				Activity: &act,
			}, nil
		case <-bgCtx.Done():
			return nil, bgCtx.Err()
		}
	case <-bgCtx.Done():
		return nil, bgCtx.Err()
	}
}

func (s *activityServer) getActivityType(ctx context.Context, uid uuid.UUID) (*activitytypecomm.ActivityType, error) {
	resp, err := s.actTypeCli.Read(ctx, &activitytypecomm.ReadActivityTypeRequest{Id: uid.String()})
	return resp.GetActivityType(), err
}

func (s *activityServer) readFromRepo(ctx context.Context, uid uuid.UUID) (*activitycomm.Activity, error) {
	act, err := s.repo.Read(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no entry found for id %s", uid.String())
	}
	return &activitycomm.Activity{
		Code:        act.Code,
		Description: act.Description,
		Name:        act.Name,
		Id:          act.ID.String(),
	}, nil
}
