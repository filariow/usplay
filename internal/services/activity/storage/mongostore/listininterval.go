package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) ListInInterval(ctx context.Context, interval storage.Interval) (storage.Activities, error) {
	filter := bson.M{
		"interval": bson.M{
			"from": bson.M{"$gte": interval.From.Unix()},
			"to":   bson.M{"$lte": interval.To.Unix()},
		},
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
