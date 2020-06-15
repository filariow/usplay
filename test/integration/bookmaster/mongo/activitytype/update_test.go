package activity_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/google/uuid"
)

func TestUpdateHappyPathIntegration(t *testing.T) {
	collection := randomString(10)
	repo, err := getRepo(collection)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	createActivityType := storage.ActivityType{
		Name:       "Test ActivityType",
		Code:       1,
		NeedsOrder: false,
	}
	id, err := repo.Create(ctx, createActivityType)
	if err != nil {
		t.Fatalf("error creating ActivityType: %v", err)
	}
	if id == nil {
		t.Fatalf("error returned id is empty")
	}

	updateActivityType := storage.ActivityType{
		ID:         id.String(),
		Code:       2,
		Name:       "Test ActivityType Updated",
		NeedsOrder: true,
	}
	if err := repo.Update(ctx, updateActivityType); err != nil {
		t.Fatalf("error updating ActivityType (%s) : %v", *id, err)
	}

	readActivityType, err := repo.Read(ctx, *id)
	if err != nil {
		t.Fatalf("error retrieving created ActivityType (%s) from repo: %v", *id, err)
	}

	if uatc, ratc := updateActivityType.Code, readActivityType.Code; uatc != ratc {
		t.Errorf("expected activity type Code (%v), obtained (%v)", uatc, ratc)
	}
	if uatn, ratn := updateActivityType.Name, readActivityType.Name; uatn != ratn {
		t.Errorf("expected activity type Name (%v), obtained (%v)", uatn, ratn)
	}
	if uatno, ratno := updateActivityType.NeedsOrder, readActivityType.NeedsOrder; uatno != ratno {
		t.Errorf("expected activity type NeedsOrder (%v), obtained (%v)", uatno, ratno)
	}
}

func TestUpdateNonExistentIntegration(t *testing.T) {
	collection := randomString(10)
	repo, err := getRepo(collection)
	if err != nil {
		t.Fatal(err)
	}

	id, ctx := uuid.New(), context.TODO()
	updateActivityType := storage.ActivityType{
		ID:         id.String(),
		Code:       2,
		Name:       "Test ActivityType Updated",
		NeedsOrder: true,
	}

	err = repo.Update(ctx, updateActivityType)
	if err == nil {
		t.Fatalf("expected error updating ActivityType (%s), none obtained", id)
	}
}
