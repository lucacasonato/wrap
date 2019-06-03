package wrap

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/imdario/mergo"
	"github.com/lucacasonato/wrap/filter"
	"github.com/lucacasonato/wrap/update"
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

// UpdateDocumentsWhere the filter matches
func (c *Collection) UpdateDocumentsWhere(filter filter.Filter, upsert bool, updates ...update.Update) error {
	var final = bson.M{}

	for _, u := range updates {
		err := mergo.Merge(&final, u)
		if err != nil {
			return err
		}
	}

	_, err := c.collection.UpdateMany(c.Database.Client.context, filter, final, options.Update().SetUpsert(upsert))
	if err != nil {
		return err
	}

	return nil
}

// DeleteDocumentsWhere the filter matches
func (c *Collection) DeleteDocumentsWhere(filter filter.Filter) error {
	_, err := c.collection.DeleteMany(c.Database.Client.context, filter)
	if err != nil {
		return err
	}

	return nil
}
