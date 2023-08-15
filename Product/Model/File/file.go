package File

import (
	"Backend/Core/Utilities/Base"
	"github.com/google/uuid"
	"time"
)

type File struct {
	Base.ResponseItem `json:"-" bson:"-"`
	URL               string
	Filename          string
	Size              int64
	CreatedAt         time.Time
	ID                string
}

func NewFile(url string, filename string, size int64) *File {
	return &File{
		URL:       url,
		Filename:  filename,
		Size:      size,
		CreatedAt: time.Now(),
		ID:        uuid.NewString(),
	}
}
