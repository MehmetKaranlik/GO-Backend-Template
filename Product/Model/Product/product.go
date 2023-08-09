package Product

import (
	"Backend/Core/Utilities/Base"
	"Backend/Product/Model/User"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Base.ResponseItem `json:"-" bson:"-"`
	Id                string
	Name              string
	Description       string
	Price             float64
	Quantity          int
	Colors            []ProductColor
	Owner             User.MetaUser
	BuyCount          int
	Rating            float64
	CreateDate        time.Time `json:"-"`
	Comments          []Comment
}

type ProductColor struct {
	ColorHex string `json:"color_hex"`
	Images   []string
}

func NewProduct(
	name string,
	description string,
	price float64,
	quantity int,
	colors []ProductColor,
	owner User.MetaUser,
) *Product {
	return &Product{
		Id:          uuid.NewString(),
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Colors:      colors,
		Owner:       owner,
		BuyCount:    0,
		Rating:      0,
		CreateDate:  time.Now(),
		Comments:    []Comment{},
	}
}

func NewProductFromResponse(body ProductCreateBody, id string, username string, profileImage string) *Product {
	return NewProduct(
		body.Name,
		body.Description,
		body.Price,
		body.Quantity,
		body.Colors,
		*User.NewMetaUser(
			id,
			username,
			profileImage,
		),
	)
}
