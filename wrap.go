package wrap

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client wraps the mongo client
type Client struct {
	client *mongo.Client
}

// Database is a database instance
type Database struct {
	ID       string
	database *mongo.Database
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
	collection *Collection
	pipes      []*bson.M
}

// Iterator to iterate over documents
type Iterator struct {
	Collection *Collection
	cursor     *mongo.Cursor
}
