package activity_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
	"github.com/google/uuid"
)

func TestDeleteHappyPathIntegration(t *testing.T) {
	collection := randomString(10)
	repo, err := getRepo(collection)
	if err != nil {
		t.Fatal(err)
	}

	// create
	ctx := context.TODO()
	createActivityType := storage.ActivityType{
		Name:       "Test ActivityType",
		Code:       1,
		NeedsOrder: false,
	}
	id, err := repo.Create(ctx, createActivityType)
	if err != nil {
		t.Fatalf("error creating new activity: %v", err)
	}
	if id == nil {
		t.Fatalf("error returned id is empty")
	}

	// read
	readActivityType, err := repo.Read(ctx, *id)
	if err != nil {
		t.Fatalf("error retrieving created ActivityType (%s) from repo: %v", *id, err)
	}

	if catc, ratc := createActivityType.Code, readActivityType.Code; catc != ratc {
		t.Errorf("expected activity type Code (%v), obtained (%v)", catc, ratc)
	}
	if catn, ratn := createActivityType.Name, readActivityType.Name; catn != ratn {
		t.Errorf("expected activity type Name (%v), obtained (%v)", catn, ratn)
	}
	if catno, ratno := createActivityType.NeedsOrder, readActivityType.NeedsOrder; catno != ratno {
		t.Errorf("expected activity type NeedsOrder (%v), obtained (%v)", catno, ratno)
	}

	// delete
	if err := repo.Delete(ctx, *id); err != nil {
		t.Fatalf("error deleting created ActivityType (%s): %v", *id, err)
	}

	// read
	readDeletedActivityType, err := repo.Read(ctx, *id)
	if err == nil {
		t.Errorf("expected error retrieving ActivityType (%s) from repo, obtained nil", *id)
	}
	if readDeletedActivityType != nil {
		t.Errorf("expected nil ActivityType, obtained %v", *readDeletedActivityType)
	}
}

func TestDeleteNotExistingIDIntegration(t *testing.T) {
	collection := randomString(10)
	repo, err := getRepo(collection)
	if err != nil {
		t.Fatal(err)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		t.Fatalf("error generating a new ID: %v", err)
	}

	// delete
	ctx := context.TODO()
	if err := repo.Delete(ctx, id); err == nil {
		t.Fatalf("error expected deleting not existing ActivityType (%s), none obtained", id)
	}
}
