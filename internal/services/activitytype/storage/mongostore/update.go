package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/errs/repoerrs"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, act storage.ActivityType) error {
	filter := bson.M{"_id": act.ID}
	data := bson.M{"$set": bson.M{
		"code":        act.Code,
		"description": act.Description,
		"name":        act.Name,
	}}
	sr, err := s.Collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	if sr.MatchedCount <= 0 {
		return &repoerrs.NotFoundError{ID: act.ID}
	}
	if sr.ModifiedCount <= 0 {
		return &repoerrs.NotUpdatedError{ID: act.ID}
	}
	return nil
}
