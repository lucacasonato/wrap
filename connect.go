package wrap

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to a mongo instance
func Connect(mongoURI string, timeout time.Duration) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:  client,
		context: context.Background(),
		timeout: timeout,
	}, nil
}
