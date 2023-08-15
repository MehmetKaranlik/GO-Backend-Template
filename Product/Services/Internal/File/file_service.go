package File

import (
	"Backend/Product/Model/File"
	"Backend/Product/Repositories"
)

type IFileService interface {
	CreateFile(file *File.File) error
}

type FileService struct {
	FileRepository Repositories.IFileRepository
}

func (self *FileService) CreateFile(file *File.File) error {
	return self.FileRepository.Create(file)
}
