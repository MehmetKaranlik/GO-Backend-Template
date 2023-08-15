package Product

import (
	"Backend/Product/Model/Product"
	"Backend/Product/Repositories"
)

type IProductService interface {
	CreateProduct(model *Product.Product) error
	CommentProduct(productId int, comment Product.Comment) error
}

type ProductService struct {
	Repository Repositories.IProductRepository
}

func (self ProductService) CreateProduct(model *Product.Product) error {
	return self.Repository.Create(model)
}

func (self ProductService) CommentProduct(productId int, comment Product.Comment) error {
	product, err := self.Repository.Read(productId)
	if err != nil {
		return err
	}
	product.Comments = append(product.Comments, comment)
	return self.Repository.Update(product)
}
