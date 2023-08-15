package Product

import (
	"Backend/Product/Services/Internal/Product"
	"net/http"
)

type FavoritesController struct {
	Service Product.IProductService
}

func (f FavoritesController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
