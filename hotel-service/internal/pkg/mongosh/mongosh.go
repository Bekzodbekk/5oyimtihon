package mongosh

import (
	"context"
	"fmt"
	"hotel-service/internal/pkg/load"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongosh struct {
	Client     mongo.Client
	Collection mongo.Collection
}

func InitDB(conf load.Config) (*Mongosh, error) {
	ctx := context.Background()

	uri := fmt.Sprintf("mongodb://%s:%d", conf.Mongosh.MongoshHost, conf.Mongosh.MongoshPort)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	coll := client.Database(conf.Mongosh.MongoshDatabase).Collection(conf.Mongosh.MongoshCollection)

	return &Mongosh{
		Client:     *client,
		Collection: *coll,
	}, nil
}
