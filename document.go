package wrap

import (
	"context"

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
	res, err := c.collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}

	return &Document{ID: res.InsertedID.(primitive.ObjectID).Hex(), Collection: c}, nil
}

// Get the contents of a document
func (d *Document) Get(data interface{}) error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	err = d.Collection.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(data)
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

	_, err = d.Collection.collection.ReplaceOne(context.Background(), bson.M{"_id": objID}, data, options.Replace().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

// Update a document using the update operators
func (d *Document) Update(update Update, upsert bool) error {
	objID, err := primitive.ObjectIDFromHex(d.ID)
	if err != nil {
		return err
	}

	_, err = d.Collection.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update, options.Update().SetUpsert(upsert))
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

	_, err = d.Collection.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
