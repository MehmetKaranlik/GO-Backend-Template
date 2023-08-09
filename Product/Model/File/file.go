package File

import (
	"Backend/Core/Utilities/Base"
	"time"
)

type File struct {
	Base.ResponseItem `json:"-" bson:"-"`
	URL               string
	Filename          string
	Size              int64
	CreatedAt         time.Time
}
