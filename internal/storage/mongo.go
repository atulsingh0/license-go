package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStorage struct {
	Collection *mongo.Collection
}

func (ms *MongoDBStorage) Save(data string) error {
	_, err := ms.Collection.InsertOne(context.Background(), data)
	return err
}

func (ms *MongoDBStorage) ReadAll() ([]string, error) {
	cur, err := ms.Collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []string
	for cur.Next(context.Background()) {
		var data string
		if err := cur.Decode(&data); err != nil {
			return nil, err
		}
		results = append(results, data)
	}
	return results, nil
}
