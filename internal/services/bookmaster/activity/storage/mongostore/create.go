package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activity/storage"
	"github.com/google/uuid"
)

func (s *mongoStore) Create(ctx context.Context, activity storage.Activity) (*uuid.UUID, error) {
	id := uuid.New()
	activity.ID = id.String()

	if _, err := s.Collection.InsertOne(ctx, activity); err != nil {
		return nil, err
	}

	return &id, nil
}
