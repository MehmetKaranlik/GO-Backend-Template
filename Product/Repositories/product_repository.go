package Repositories

import (
	"Backend/Core/Utilities/Base"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Model/Product"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IProductRepository interface {
	Base.IRepository[Product.Product]
}

type ProductRepository struct {
	Ref *Mongo.MongoCollectionRef
}

func (self ProductRepository) Create(model *Product.Product) error {
	ctx := context.Background()
	_, err := self.Ref.Ref.InsertOne(ctx, model, options.InsertOne())
	return err
}

func (self ProductRepository) Read(id int) (*Product.Product, error) {
	filter := bson.M{"id": id}
	ctx := context.Background()
	var result Product.Product
	err := self.Ref.Ref.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	return &result, err
}

func (self ProductRepository) Update(model *Product.Product) error {
	filter := bson.M{"id": model.Id}
	ctx := context.Background()
	_, err := self.Ref.Ref.UpdateOne(ctx, filter, model, options.Update())
	return err
}

func (self ProductRepository) Delete(id int) error {
	filter := bson.M{"id": id}
	ctx := context.Background()
	_, err := self.Ref.Ref.DeleteOne(ctx, filter, options.Delete())
	return err
}

func (self ProductRepository) Aggregate(pipeline []interface{}) (*[]Product.Product, error) {
	ctx := context.Background()
	cursor, err := self.Ref.Ref.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		return nil, err
	}
	var result []Product.Product
	err = cursor.All(ctx, &result)
	return &result, err
}
