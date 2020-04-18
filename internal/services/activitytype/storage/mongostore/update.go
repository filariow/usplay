package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/repoerrors"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, act storage.ActivityType) error {
	filter := bson.M{"_id": act.ID}

	sr, err := s.Collection.UpdateOne(ctx, filter, act)
	if err != nil {
		return err
	}

	if sr.MatchedCount <= 0 {
		return &repoerrors.NotFoundError{ID: act.ID}
	}
	if sr.ModifiedCount <= 0 {
		return &repoerrors.NotUpdatedError{ID: act.ID}
	}
	return nil
}
