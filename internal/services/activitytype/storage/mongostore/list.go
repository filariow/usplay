package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) List(ctx context.Context, ids []uuid.UUID) (storage.Activities, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []storage.ActivityType
	if err := cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
