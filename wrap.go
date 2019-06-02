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

// Update fields
type Update bson.M
