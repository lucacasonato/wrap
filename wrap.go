package wrap

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client wraps the mongo client
type Client struct {
	client  *mongo.Client
	context context.Context
}

// Database is a database instance
type Database struct {
	ID       string
	database *mongo.Database
	Client   *Client
}

// Collection is a collection on the database
type Collection struct {
	ID         string
	collection *mongo.Collection
	Database   *Database
}

// Document is a document in a collection
type Document struct {
	ID         string
	Collection *Collection
}

// DocumentData is the data in a document
type DocumentData struct {
	Document *Document
	result   *mongo.SingleResult
}

// CollectionQuery is a filtered abstraction of a group of documents
type CollectionQuery struct {
	Collection *Collection
	pipes      []*bson.M
}

// Iterator to iterate over documents
type Iterator struct {
	Collection *Collection
	cursor     *mongo.Cursor
}

// BulkCollection is a collection which is used for bulk writing
type BulkCollection struct {
	ID         string
	Collection *Collection
	models     []mongo.WriteModel
}

// BulkDocument is a document which is used for bulk writing
type BulkDocument struct {
	ID         string
	Collection *BulkCollection
}
