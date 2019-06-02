package wrap_test

import (
	"testing"

	"github.com/lucacasonato/wrap"
)

func createDatabase() (*wrap.Database, error) {
	client, err := connect()
	if err != nil {
		return nil, err
	}

	database := client.Database("testing")

	return database, nil
}

func TestDatabase(t *testing.T) {
	database, err := createDatabase()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(database)
}

func TestDatabaseDelete(t *testing.T) {
	database, err := createDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
