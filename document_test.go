package wrap_test

import (
	"testing"

	"github.com/lucacasonato/wrap/update"
)

func TestCollectionAddUpdate(t *testing.T) {
	collection, err := createCollection()
	if err != nil {
		t.Fatal(err)
	}

	doc, err := collection.Add(map[string]interface{}{
		"name": "the red fish",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(doc)

	err = doc.Update(update.Set("name", "The red fish."), true)
	if err != nil {
		t.Fatal(err)
	}

	err = collection.Database.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDocumentSetGet(t *testing.T) {
	collection, err := createCollection()
	if err != nil {
		t.Fatal(err)
	}

	redFish := collection.Document("0123456789abcdef01234567")

	err = redFish.Set(map[string]interface{}{
		"name": "the red fish 2",
	})
	if err != nil {
		t.Fatal(err)
	}

	var fishData map[string]interface{}

	doc, err := redFish.Get()
	if err != nil {
		t.Fatal(err)
	}

	err = doc.DataTo(&fishData)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(fishData)

	err = collection.Database.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
