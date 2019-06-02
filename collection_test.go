package wrap_test

import (
	"testing"

	"github.com/lucacasonato/wrap"
)

func createCollection() (*wrap.Collection, error) {
	database, err := createDatabase()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("fish")

	return collection, nil
}

func TestCollection(t *testing.T) {
	collection, err := createCollection()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(collection)
}

func TestCollectionDelete(t *testing.T) {
	collection, err := createCollection()
	if err != nil {
		t.Fatal(err)
	}

	err = collection.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
