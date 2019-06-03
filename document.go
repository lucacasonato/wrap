package wrap

import (
	"github.com/imdario/mergo"
	"github.com/lucacasonato/wrap/update"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dataInterface interface{}

type documentData struct {
	dataInterface
	ID primitive.ObjectID `bson:"_id"`
}

// Document get a single document from a collection
func (c *Collection) Document(id string) *Document {
	return &Document{ID: id, Collection: c}
}

// Add a document with a certain value
func (c *Collection) Add(data interface{}) (*Document, error) {
	res, err := c.collection.InsertOne(c.Database.Client.context, data)
	if err != nil {
		return nil, err
	}

	return &Document{ID: res.InsertedID.(primitive.ObjectID).Hex(), Collection: c}, nil
}

// Get the contents of a document
func (d *Document) Get() (*DocumentData, error) {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return nil, err
	}

	return &DocumentData{
		Document: d,
		result:   d.Collection.collection.FindOne(d.Collection.Database.Client.context, bson.M{"_id": objID}),
	}, nil
}

// Data decodes some data and returns an interface
func (d *DocumentData) Data() (interface{}, error) {
	var data interface{}

	err := d.result.Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DataTo decodes some data into an interface
func (d *DocumentData) DataTo(data interface{}) error {
	err := d.result.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// Set a document to a certain value
func (d *Document) Set(data interface{}) error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	_, err = d.Collection.collection.ReplaceOne(d.Collection.Database.Client.context, bson.M{"_id": objID}, data, options.Replace().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

// Update a document using the update operators
func (d *Document) Update(upsert bool, updates ...update.Update) error {
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

	_, err = d.Collection.collection.UpdateOne(d.Collection.Database.Client.context, bson.M{"_id": objID}, final, options.Update().SetUpsert(upsert))
	if err != nil {
		return err
	}

	return nil
}

// Delete a document from a collection
func (d *Document) Delete() error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	_, err = d.Collection.collection.DeleteOne(d.Collection.Database.Client.context, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
