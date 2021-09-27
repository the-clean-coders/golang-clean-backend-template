package drivers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDriver struct {
	db         *mongo.Database
	collection string
}

func NewMongoDriver(db *mongo.Database) *MongoDriver {
	return &MongoDriver{
		db:         db,
		collection: "examples",
	}
}

func (md *MongoDriver) GetExample(ctx context.Context, id string) (*ExampleModel, error) {
    example := new(ExampleModel)

	err := md.db.Collection(md.collection).FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(example)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return example, nil
}
