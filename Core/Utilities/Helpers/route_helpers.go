package Helpers

import (
	"Backend/Core/App/Router"
	"Backend/Core/Middleware"
	"Backend/Product/Init/Databases/Mongo"
	"net/http"
)

func InitializeRoutes(routes []Router.Route) {
	for _, route := range routes {
		registerRoutes(route.Pattern, route.Handler)
	}
}

func registerRoutes(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func MakeJsonAuthorizedRoute(pattern string, database *Mongo.MongoDatabaseRef, handler http.Handler) Router.Route {
	return Router.Route{
		Pattern: pattern,
		Handler: &Middleware.AuthenticationMiddleware{
			Connection: database.Users,
			Next:       &Middleware.JSONContentTypeMiddleware{Next: handler},
		},
	}
}

func MakeJsonRoute(pattern string, handler http.Handler) Router.Route {
	return Router.Route{
		Pattern: pattern,
		Handler: &Middleware.JSONContentTypeMiddleware{Next: handler},
	}
}
