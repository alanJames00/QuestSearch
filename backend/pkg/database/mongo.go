// mongodb setup and connections
package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to mongodb, an error from this function should panic the code
func ConnectToMongo(uri string) (*mongo.Client, error) {
	// 10 seconds context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("connected to mongodb", uri)
	
	return client, nil
}
