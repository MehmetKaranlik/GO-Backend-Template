package Middleware

import (
	"Backend/Core/Constants/Enums/Headers"
	"net/http"
)

type JSONContentTypeMiddleware struct {
	Next http.Handler
}

func (self *JSONContentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if self.Next == nil {
		self.Next = http.DefaultServeMux
	}
	w.Header().Set(Headers.ContentTypeKey, Headers.JSON.Value())
	self.Next.ServeHTTP(w, r)
}
