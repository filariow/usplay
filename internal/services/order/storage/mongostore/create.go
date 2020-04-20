package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/google/uuid"
)

func (s *mongoStore) Create(ctx context.Context, order storage.Order) (*uuid.UUID, error) {
	id := uuid.New()
	order.ID = id.String()

	if _, err := s.Collection.InsertOne(ctx, order); err != nil {
		return nil, err
	}

	return &id, nil
}
