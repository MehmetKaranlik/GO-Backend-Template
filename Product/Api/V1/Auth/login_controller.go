package Auth

import (
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Services/Internal/Auth"
	"net/http"
)

type LoginController struct {
	Service Auth.AuthService
}

func (self LoginController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	user, err := self.Service.Login(request.Body)
	if err != nil {
		Responses.InvokeBadRequest(writer, err.Error())
		return
	}

	Responses.InvokeSuccess(writer, user)

}
