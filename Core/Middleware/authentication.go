package Middleware

import (
	JWT "Backend/Core/Utilities/Jwt"
	"Backend/Core/Utilities/Responses"
	"Backend/Product/Init/Databases/Mongo"

	"net/http"
	"strings"
)

const (
	authorizationKey = "Authorization"
	separator        = " "
	tokenNotFound    = "Token not found"
	tokenIsNotValid  = "Token is not valid"
)

type AuthenticationMiddleware struct {
	Next       http.Handler
	Connection *Mongo.MongoCollectionRef
}

func (h *AuthenticationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Next == nil {
		h.Next = http.DefaultServeMux
	}

	tokenString := validateToken(w, r)

	if tokenString == "" {
		Responses.InvokeUnAuthorized(w, tokenNotFound)
		return
	}

	id, err := JWT.ParseAccessToken(tokenString)
	if err != nil || id == "" {
		Responses.InvokeUnAuthorized(w, tokenIsNotValid)
		return
	}

	h.Next.ServeHTTP(w, r)

}

func validateToken(w http.ResponseWriter, r *http.Request) string {
	token := extractToken(r)
	if token == "" {
		return ""
	}
	return token
}

func extractToken(r *http.Request) string {
	tokenString := r.Header.Get(authorizationKey)
	tokenParts := strings.Split(tokenString, separator)
	if len(tokenParts) < 2 {
		return ""
	}
	return string(tokenParts[len(tokenParts)-1])

}
