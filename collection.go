package wrap

import (
	"context"
)

// Collection in the database by id
func (d *Database) Collection(id string) *Collection {
	collection := d.database.Collection(id)

	return &Collection{ID: id, collection: collection, Database: d}
}

// Delete a collection
func (c *Collection) Delete() error {
	return c.collection.Drop(context.Background())
}
