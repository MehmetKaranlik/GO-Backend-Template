package Middlewares

import (
	"Backend/Core/Constants/Keys/ContextKeys"
	"Backend/Core/Utilities/Methods"
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Enums"
	"Backend/Product/Model/User"
	"net/http"
)

const (
	roleBasedUnAuthorizedError = "You are not allowed to access this resource"
)

type RoleCheckerMiddleware struct {
	Next         http.Handler
	AllowedRoles []Enums.UserRole
}

func (r RoleCheckerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(ContextKeys.UserKey).(User.User)

	if !Methods.Contains(r.AllowedRoles, user.Role) {
		Responses.InvokeUnAuthorized(writer, roleBasedUnAuthorizedError)
		return
	}

	if r.Next == nil {
		r.Next = http.DefaultServeMux
	}
	r.Next.ServeHTTP(writer, request)
}
