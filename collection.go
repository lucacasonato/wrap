package wrap

import (
	"context"

	"github.com/lucacasonato/wrap/filter"
)

// Collection in the database by id
func (d *Database) Collection(id string) *Collection {
	collection := d.database.Collection(id)

	return &Collection{ID: id, collection: collection, Database: d}
}

// Where returns an abstract of the collection of fields that match the filter
func (c *Collection) Where(filter filter.Filter) *CollectionQuery {
	return &CollectionQuery{
		collection: c,
		filter:     filter,
	}
}

// DocumentIterator gives you an iterator to loop over the documents
func (c *Collection) DocumentIterator() (*Iterator, error) {
	cursor, err := c.collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &Iterator{
		cursor: cursor,
	}, nil
}

// Delete a collection
func (c *Collection) Delete() error {
	return c.collection.Drop(context.Background())
}
