package Product

import (
	"Backend/Core/Constants/Keys/ContextKeys"
	"Backend/Core/Utilities/Responses"
	"Backend/Core/Utilities/Validation"
	"Backend/Product/Model/Product"
	"Backend/Product/Model/User"
	service "Backend/Product/Services/Internal/Product"
	"net/http"
)

const (
	successMessage = "Product created successfully"
)

type ProductCreateController struct {
	Service service.IProductService
}

func (p ProductCreateController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body Product.ProductCreateBody

	if err := Validation.NewCustomJsonDecoder(request.Body).Decode(&body); err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	user := request.Context().Value(ContextKeys.UserKey).(*User.User)

	product := Product.NewProductFromResponse(body, user.Id, user.Username, user.ProfileImage)

	if err := p.Service.CreateProduct(product); err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	Responses.InvokeSuccess(writer, successMessage)

}
