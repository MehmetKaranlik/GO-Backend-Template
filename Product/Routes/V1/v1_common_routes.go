package V1

import (
	"Backend/Core/App/Router"
	"Backend/Core/Middleware"
	"Backend/Core/Utilities/Helpers"
	"Backend/Product/Api/V1/Auth"
	"Backend/Product/Enums"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Middlewares"
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
		self.createSafeRoute(
			ROUTE_PREFIX+"/auth/login",
			[]string{http.MethodPost},
			[]Enums.UserRole{Enums.Normal, Enums.Seller, Enums.Admin},
			Auth.LoginController{
				Service: Auth2.AuthService{
					UserRepository: Repositories.UserRepository{
						Ref: Mongo.MongoConnectionItem.Main.Users,
					},
				},
			},
		),
	)

	routes = append(
		routes,
		self.createSafeRoute(
			ROUTE_PREFIX+"/auth/register",
			[]string{http.MethodPost},
			[]Enums.UserRole{Enums.Normal, Enums.Seller, Enums.Admin},
			&Auth.RegisterController{
				Service: Auth2.AuthService{
					UserRepository: Repositories.UserRepository{
						Ref: Mongo.MongoConnectionItem.Main.Users,
					},
				},
			},
		),
	)

	return routes
}

func (self *V1Routes) createSafeRoute(
	Pattern string,
	AllowedMethods []string,
	AllowedRoles []Enums.UserRole,
	Next http.Handler,
) Router.Route {
	return Helpers.MakeJsonAuthorizedRoute(
		Pattern,
		Mongo.MongoConnectionItem.Main,
		&Middleware.MethodCheckerMiddleware{
			AllowedMethods: AllowedMethods,
			Next: &Middlewares.RoleCheckerMiddleware{
				AllowedRoles: AllowedRoles,
				Next:         Next,
			},
		},
	)
}
