package Product

import (
	"Backend/Core/Constants/Keys/ContextKeys"
	"Backend/Core/Constants/Keys/ErrorKeys"
	"Backend/Core/Utilities/Helpers"
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Model/User"
	"Backend/Product/Services/Internal/Feed"
	"net/http"
)

type FavoritesController struct {
	Service Feed.IFeedService
}

func (f FavoritesController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(ContextKeys.UserKey).(*User.User)
	page, limit, err := Helpers.StripPaginationFromUrl(request)
	if err != nil {
		Responses.InvokeBadRequest(writer, ErrorKeys.InvalidPagination)
		return
	}
	products, err := f.Service.Favorites(int(page), int(limit), user)
	if err != nil {
		Responses.InvokeInternalServerError(writer, err.Error())
		return
	}
	Responses.InvokeSuccess(writer, products)
}
