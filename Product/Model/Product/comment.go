package Product

import (
	"Backend/Product/Model/User"
	"time"
)

type Comment struct {
	Id         string
	Content    string
	Owner      User.MetaUser
	Rating     int
	CreateDate time.Time
}
