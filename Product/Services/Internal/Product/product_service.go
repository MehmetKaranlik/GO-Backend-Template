package Product

import (
	"Backend/Product/Model/Product"
	"Backend/Product/Repositories"
)

type IProductService interface {
	CreateProduct(model *Product.Product) error
	CommentProduct(productId string, comment Product.Comment) error
}

type ProductService struct {
	Repository Repositories.IProductRepository
}

func (self ProductService) CreateProduct(model *Product.Product) error {
	return self.Repository.Create(model)
}

func (self ProductService) CommentProduct(productId string, comment Product.Comment) error {
	//TODO implement me
	panic("implement me")
}
