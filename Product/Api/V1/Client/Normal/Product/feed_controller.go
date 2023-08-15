package Product

import "Backend/Product/Services/Internal/Product"

type FeedController struct {
	Service Product.IProductService
}
