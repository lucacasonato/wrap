package wrap_test

import (
	"testing"

	"github.com/lucacasonato/wrap"
)

func TestBulkWrite(t *testing.T) {
	collection, err := createCollection()
	if err != nil {
		t.Fatal(err)
	}

	err = collection.Bulk(func (c *wrap.BulkCollection) error {
		c.Add(map[string]interface{}{
			"name": "the red fish",
		})
		if err != nil {
			return err
		}

		blueFish := c.Document("0123456789abcdef12345678")

		err = blueFish.Set(map[string]interface{}{
			"name": "the blue fish",
		})
		if err != nil {
			return err
		}

		return nil
	}, false)
	if err != nil {
		t.Fatal(err)
	}
}
