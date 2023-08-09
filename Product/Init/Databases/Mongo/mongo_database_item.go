package Mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "main"
	userCollection = "users"
	products       = "products"
	gridFSBucket   = "fs"
)

var MongoConnectionItem MongoClusterRef

type MongoClusterRef struct {
	Client *mongo.Client
	Main   *MongoDatabaseRef
}

type MongoDatabaseRef struct {
	Ref      *mongo.Database
	Users    *MongoCollectionRef
	Products *MongoCollectionRef
}

type MongoCollectionRef struct {
	Ref *mongo.Collection
}

func (r *MongoClusterRef) MapFromConnection(c *mongo.Client) {
	r.Client = c
	r.Main = &MongoDatabaseRef{
		Ref:      c.Database(dbName),
		Users:    &MongoCollectionRef{Ref: c.Database(dbName).Collection(userCollection)},
		Products: &MongoCollectionRef{Ref: c.Database(dbName).Collection(products)},
	}
}
