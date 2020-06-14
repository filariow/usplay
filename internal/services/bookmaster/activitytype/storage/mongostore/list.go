package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) List(ctx context.Context, ids []uuid.UUID) (storage.ActivityTypes, error) {
	idsString := make([]string, len(ids))
	for i, id := range ids {
		idsString[i] = id.String()
	}

	filter := bson.M{}
	if len(idsString) > 0 {
		filter = bson.M{"_id": bson.M{"$in": idsString}}
	}

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
