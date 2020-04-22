package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) ListInInterval(ctx context.Context, from, to *timestamp.Timestamp) (storage.Activities, error) {
	filter := bson.M{
		"interval": bson.M{
			"from": bson.M{"$gte": from.GetSeconds()},
			"to":   bson.M{"$lte": to.GetSeconds()},
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
