package mongosh

import (
	"context"
	"fmt"
	"user-service/internal/pkg/load"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func InitDB(conf load.Config)(*Mongo, error){
	ctx := context.Background()
	
	uri := fmt.Sprintf("mongodb://%s:%d", conf.Mongo.MongoHost, conf.Mongo.MongoPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil{
		return nil, err
	}

	coll := client.Database(conf.Mongo.MongoDatabase).Collection(conf.Mongo.MongoCollection)

	return &Mongo{
		Client: client,
		Collection: coll,
	}, nil
}
