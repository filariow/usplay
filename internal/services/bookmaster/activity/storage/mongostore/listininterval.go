package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) ListInInterval(ctx context.Context, interval storage.Interval) (storage.Activities, error) {
	filter := bson.M{
		"interval.from": bson.M{"$gte": interval.From},
		"interval.to":   bson.M{"$lte": interval.To},
	}

	cursor, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []storage.Activity
	if err := cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
