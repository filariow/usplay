package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/google/uuid"
)

func (s *mongoStore) Create(ctx context.Context, act storage.ActivityType) (*uuid.UUID, error) {
	id := uuid.New()
	act.ID = id.String()

	if _, err := s.Collection.InsertOne(ctx, act); err != nil {
		return nil, err
	}

	return &id, nil
}
