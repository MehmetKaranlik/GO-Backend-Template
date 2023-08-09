package V1

import (
	"Backend/Core/App/Router"
	"Backend/Core/Middleware"
	"Backend/Core/Utilities/Helpers"
	"Backend/Product/Api/V1/Auth"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Repositories"
	Auth2 "Backend/Product/Services/Internal/Auth"
	"net/http"
)

const (
	ROUTE_PREFIX = "/api/v1"
)

type V1Routes struct {
}

func (self V1Routes) MakeRoutes() []Router.Route {
	routes := make([]Router.Route, 0)

	// Auth
	routes = append(
		routes,
		Helpers.MakeJsonRoute(
			ROUTE_PREFIX+"/auth/login",
			&Middleware.MethodCheckerMiddleware{
				Next: &Auth.LoginController{
					Service: Auth2.AuthService{
						UserRepository: Repositories.UserRepository{
							Ref: Mongo.MongoConnectionItem.Main.Users,
						},
					},
				},
				AllowedMethods: []string{http.MethodPost},
			},
		),
	)

	routes = append(
		routes,
		Helpers.MakeJsonRoute(
			ROUTE_PREFIX+"/auth/register",
			&Middleware.MethodCheckerMiddleware{
				Next: &Auth.RegisterController{
					Service: Auth2.AuthService{
						UserRepository: Repositories.UserRepository{
							Ref: Mongo.MongoConnectionItem.Main.Users,
						},
					},
				},
				AllowedMethods: []string{http.MethodPost},
			},
		),
	)

	return routes
}
