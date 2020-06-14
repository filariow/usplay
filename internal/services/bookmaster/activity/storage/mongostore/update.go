package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/errs/repoerrs"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, activity storage.Activity) error {
	filter := bson.M{"_id": activity.ID}
	data := bson.M{"$set": bson.M{
		"interval":        activity.Period,
		"activitytype_id": activity.ActivityTypeID,
		"order_id":        activity.OrderID,
	}}
	sr, err := s.Collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	if sr.MatchedCount <= 0 {
		return &repoerrs.NotFoundError{ID: activity.ID}
	}
	if sr.ModifiedCount <= 0 {
		return &repoerrs.NotUpdatedError{ID: activity.ID}
	}
	return nil
}
