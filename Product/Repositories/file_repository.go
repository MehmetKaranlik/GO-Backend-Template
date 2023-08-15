package Repositories

import (
	"Backend/Core/Utilities/Base"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Model/File"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IFileRepository interface {
	Base.IRepository[File.File]
}

type FileRepository struct {
	CollectionRef *Mongo.MongoCollectionRef
}

func (self FileRepository) Create(file *File.File) error {
	ctx := context.Background()
	options := options.InsertOne()
	_, err := self.CollectionRef.Ref.InsertOne(ctx, file, options)
	return err
}

func (self FileRepository) Read(id string) *File.File {
	filter := bson.M{"id": id}
	ctx := context.Background()
	var result File.File
	if err := self.CollectionRef.Ref.FindOne(ctx, filter, options.FindOne()).Decode(&result); err != nil {
		return nil
	}
	return &result
}

func (self FileRepository) Update(file *File.File) error {
	filter := bson.M{"id": file.ID}
	ctx := context.Background()
	_, err := self.CollectionRef.Ref.UpdateOne(ctx, filter, file, options.Update())
	return err
}

func (self FileRepository) Delete(id string) error {
	filter := bson.M{"id": id}
	ctx := context.Background()
	_, err := self.CollectionRef.Ref.DeleteOne(ctx, filter, options.Delete())
	return err
}

func (self FileRepository) Aggregate(pipeline []interface{}) (*[]File.File, error) {
	ctx := context.Background()
	cursor, err := self.CollectionRef.Ref.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		return nil, err
	}
	var result []File.File
	err = cursor.All(ctx, &result)
	return &result, err
}
