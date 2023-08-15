package Product

import (
	"Backend/Core/Constants/Keys/ErrorKeys"
	"Backend/Core/Utilities/Helpers"
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Services/Internal/Feed"
	"net/http"
)

type FeedController struct {
	Service Feed.IFeedService
}

func (self FeedController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	page, limit, err := Helpers.StripPaginationFromUrl(request)
	if err != nil {
		Responses.InvokeBadRequest(writer, ErrorKeys.InvalidPagination)
		return
	}
	products, err := self.Service.Feed(int(page), int(limit))
	if err != nil {
		Responses.InvokeInternalServerError(writer, err.Error())
		return
	}
	Responses.InvokeSuccess(writer, products)
}
