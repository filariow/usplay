package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/errs/repoerrs"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongoStore) Update(ctx context.Context, order storage.Order) error {
	filter := bson.M{"_id": order.ID}
	data := bson.M{"$set": bson.M{
		"code":        order.Code,
		"description": order.Description,
		"name":        order.Name,
	}}
	sr, err := s.Collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}

	if sr.MatchedCount <= 0 {
		return &repoerrs.NotFoundError{ID: order.ID}
	}
	if sr.ModifiedCount <= 0 {
		return &repoerrs.NotUpdatedError{ID: order.ID}
	}
	return nil
}
