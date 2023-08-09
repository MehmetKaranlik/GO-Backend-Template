package Product

type ProductCreateBody struct {
	Name        string         `json:"name" validate:"required,min=6,max=50"`
	Description string         `json:"description" validate:"required,min=6,max=500"`
	Price       float64        `json:"price" validate:"required,min=0,max=1000000000"`
	Quantity    int            `json:"quantity" validate:"required,min=0,max=1000000000"`
	Colors      []ProductColor `json:"colors" validate:"required,min=1,max=1000000000"`
}
