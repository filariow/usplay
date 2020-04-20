package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/pkg/errs/repoerrs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id.String()}

	res, err := s.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount <= 0 {
		return &repoerrs.NotFoundError{ID: id.String()}
	}

	return nil
}
