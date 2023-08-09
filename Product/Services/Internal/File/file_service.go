package File

import (
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Model/File"
	"Backend/Product/Repositories"
)

type IFileService interface {
	CreateFile(file *File.File) error
}

type FileService struct {
	FileRef        *Mongo.MongoCollectionRef
	FileRepository Repositories.IFileRepository
}

func (self *FileService) CreateFile(file *File.File) error {
	return nil
}
