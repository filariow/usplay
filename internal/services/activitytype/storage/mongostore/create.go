package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/google/uuid"
)

func (s *mongoStore) Create(ctx context.Context, act storage.ActivityType) (*uuid.UUID, error) {
	act.ID = uuid.New()

	if _, err := s.Collection.InsertOne(ctx, act); err != nil {
		return nil, err
	}

	return &act.ID, nil
}
