package wrap

import (
	"github.com/imdario/mergo"
	"github.com/lucacasonato/wrap/filter"
	"github.com/lucacasonato/wrap/update"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Bulk is used to do bulk writes
func (c *Collection) Bulk(run func(collection *BulkCollection) error, ordered bool) error {
	bulkCollection := &BulkCollection{
		Collection: c,
		models:     []mongo.WriteModel{},
	}

	err := run(bulkCollection)
	if err != nil {
		return err
	}

	_, err = c.collection.BulkWrite(bulkCollection.Collection.Database.Client.ctx(), bulkCollection.models, options.BulkWrite().SetOrdered(ordered))
	if err != nil {
		return err
	}

	return nil
}

// UpdateDocumentsWhere the filter matches
func (c *BulkCollection) UpdateDocumentsWhere(filter filter.Filter, upsert bool, updates ...update.Update) error {
	var final = bson.M{}

	for _, u := range updates {
		err := mergo.Merge(&final, u)
		if err != nil {
			return err
		}
	}

	c.models = append(c.models, mongo.NewUpdateManyModel().SetFilter(filter).SetUpdate(final).SetUpsert(upsert))

	return nil
}

// DeleteDocumentsWhere the filter matches
func (c *BulkCollection) DeleteDocumentsWhere(filter filter.Filter) {
	c.models = append(c.models, mongo.NewDeleteManyModel().SetFilter(filter))
}

// Document get a single document from a collection
func (c *BulkCollection) Document(id string) *BulkDocument {
	return &BulkDocument{
		ID:         id,
		Collection: c,
	}
}

// Add a document with a certain value
func (c *BulkCollection) Add(data interface{}) {
	c.models = append(c.models, mongo.NewInsertOneModel().SetDocument(data))
}

// Set a document to a certain value
func (d *BulkDocument) Set(data interface{}) error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	d.Collection.models = append(d.Collection.models, mongo.NewReplaceOneModel().SetFilter(bson.M{"_id": objID}).SetReplacement(data).SetUpsert(true))

	return nil
}

// Update a document using the update operators
func (d *BulkDocument) Update(upsert bool, updates ...update.Update) error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	var final = bson.M{}

	for _, update := range updates {
		err := mergo.Merge(&final, update)
		if err != nil {
			return err
		}
	}

	d.Collection.models = append(d.Collection.models, mongo.NewUpdateOneModel().SetFilter(bson.M{"_id": objID}).SetUpdate(final).SetUpsert(upsert))

	return nil
}

// Delete a document from a collection
func (d *BulkDocument) Delete() error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	d.Collection.models = append(d.Collection.models, mongo.NewDeleteOneModel().SetFilter(bson.M{"_id": objID}))

	return nil
}
