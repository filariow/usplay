package mongostore

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongoStore) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	filter := bson.M{"_id": id.String()}

	cursor, err := s.Collection.Find(ctx, filter, options.Find().SetLimit(1))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	res := true
	return &res, nil
}
