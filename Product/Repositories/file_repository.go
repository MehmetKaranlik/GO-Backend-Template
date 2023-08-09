package Repositories

import (
	"Backend/Core/Utilities/Base"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Model/File"
)

type IFileRepository interface {
	Base.IRepository[File.File]
}

type FileRepository struct {
	CollectionRef *Mongo.MongoCollectionRef
}

func (f FileRepository) Create(file *File.File) error {
	//TODO implement me
	panic("implement me")
}

func (f FileRepository) Read(id string) *File.File {
	//TODO implement me
	panic("implement me")
}

func (f FileRepository) Update(file *File.File) error {
	//TODO implement me
	panic("implement me")
}

func (f FileRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (f FileRepository) Aggregate(pipeline []interface{}) (*[]File.File, error) {
	//TODO implement me
	panic("implement me")
}
