package Router

import (
	middleware2 "Backend/Core/Middleware"
	"net/http"
)

func InitializeRoutes(routes []Route) {
	for _, route := range routes {
		registerRoutes(route.Pattern, route.Handler)
	}
}

func registerRoutes(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func MakeAuthorizedRoute(pattern string, handler http.Handler) Route {
	return Route{
		Pattern: pattern,
		Handler: &middleware2.AuthenticationMiddleware{
			Next: handler,
		},
	}
}

func MakeJsonAuthorizedRoute(pattern string, handler http.Handler) Route {
	return Route{
		Pattern: pattern,
		Handler: &middleware2.AuthenticationMiddleware{
			Next: &middleware2.JSONContentTypeMiddleware{Next: handler},
		},
	}
}

func MakeJsonRoute(pattern string, handler http.Handler) Route {
	return Route{
		Pattern: pattern,
		Handler: &middleware2.JSONContentTypeMiddleware{Next: handler},
	}
}
