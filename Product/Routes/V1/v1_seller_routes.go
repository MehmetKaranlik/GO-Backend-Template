package V1

import (
	"Backend/Core/App/Router"
	"Backend/Core/Middleware"
	"Backend/Core/Utilities/Helpers"
	"Backend/Product/Api/V1/Client/Seller/File"
	"Backend/Product/Api/V1/Client/Seller/Product"
	"Backend/Product/Enums"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Middlewares"
	"Backend/Product/Repositories"
	"Backend/Product/Services/External/CDN"
	Product2 "Backend/Product/Services/Internal/Product"
	"net/http"
)

type V1SellerRoutes struct {
}

func (self V1SellerRoutes) MakeRoutes() []Router.Route {
	routes := make([]Router.Route, 0)

	routes = append(
		routes,
		self.CreateSafeRoute(
			ROUTE_PREFIX+"/seller/upload-file",
			[]string{http.MethodPost},
			[]Enums.UserRole{Enums.Seller},
			&File.UploadController{
				CDN: &CDN.AWSCDNService{},
			},
		),
	)

	routes = append(
		routes,
		self.CreateSafeRoute(
			ROUTE_PREFIX+"/seller/product/create",
			[]string{http.MethodPost},
			[]Enums.UserRole{Enums.Seller},
			&Product.ProductCreateController{
				Service: Product2.ProductService{
					Repository: Repositories.ProductRepository{
						Ref: Mongo.MongoConnectionItem.Main.Products,
					},
				},
			},
		),
	)
	return routes
}

func (self *V1SellerRoutes) CreateSafeRoute(
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
