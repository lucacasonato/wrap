package wrap

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lucacasonato/wrap/filter"
)

// Collection in the database by id
func (d *Database) Collection(id string) *Collection {
	collection := d.database.Collection(id)

	return &Collection{ID: id, collection: collection, Database: d}
}

// Where returns an abstract of the collection of documents that match the filter
func (c *Collection) Where(filter filter.Filter) *CollectionQuery {
	return &CollectionQuery{
		Collection: c,
		pipes: []*bson.M{&bson.M{
			"$match": filter,
		}},
	}
}

// All returns an abstract of the collection of all documents
func (c *Collection) All() *CollectionQuery {
	return &CollectionQuery{
		Collection: c,
		pipes:      []*bson.M{},
	}
}

// CreateIndex for a single or group of fields
func (c *Collection) CreateIndex(fields map[string]Index) error {
	i := bson.M{}

	for field, index := range fields {
		i[field] = index
	}

	_, err := c.collection.Indexes().CreateOne(c.Database.Client.context, mongo.IndexModel{
		Keys: i,
	})
	if err != nil {
		return err
	}

	return nil
}

// Delete a collection
func (c *Collection) Delete() error {
	return c.collection.Drop(c.Database.Client.context)
}
