package Middleware

import (
	"Backend/Core/Constants/Keys/ErrorKeys"
	"Backend/Core/Utilities/Methods"
	"Backend/Core/Utilities/Responses"
	"net/http"
)

type MethodCheckerMiddleware struct {
	Next           http.Handler
	AllowedMethods []string
}

func (self *MethodCheckerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if !Methods.Contains(self.AllowedMethods, request.Method) {
		Responses.InvokeMethodNotAllowed(writer, ErrorKeys.InvalidMethod)
		return
	}

	if self.Next == nil {
		self.Next = http.DefaultServeMux
	}

	self.Next.ServeHTTP(writer, request)

}
