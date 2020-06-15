package activity_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/bookmaster/activitytype/storage"
)

func TestCreate(t *testing.T) {
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
		t.Fatalf("error creating new activity: %v", err)
	}
	if id == nil {
		t.Fatalf("error returned id is empty")
	}

	readActivityType, err := repo.Read(ctx, *id)
	if err != nil {
		t.Fatalf("error retrieving created activity from repo: %v", err)
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
}
