package Feed

import (
	"Backend/Product/Model/Product"
	"Backend/Product/Model/User"
	"Backend/Product/Repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type IFeedService interface {
	Feed(page int, limit int) (*[]Product.Product, error)
	Favorites(page int, limit int, user *User.User) (*[]Product.Product, error)
}

type FeedService struct {
	Repository Repositories.IProductRepository
}

func (self FeedService) Feed(page int, limit int) (*[]Product.Product, error) {
	pipeline := feedPipeline(page, limit)
	return self.Repository.Aggregate(pipeline)
}

func (self FeedService) Favorites(page int, limit int, user *User.User) (*[]Product.Product, error) {
	pipeline := favoritesPipeline(page, limit, user)
	return self.Repository.Aggregate(pipeline)
}

func feedPipeline(page int, limit int) bson.A {
	return bson.A{
		bson.M{"$sort": bson.M{"createAt": -1}},
		bson.M{"$skip": (page - 1) * limit},
		bson.M{"$limit": limit},
	}
}

func favoritesPipeline(page int, limit int, user *User.User) bson.A {
	return bson.A{
		bson.M{"$match": bson.M{"id": bson.M{"$in": user.Favorites}}},
		bson.M{"$sort": bson.M{"createAt": -1}},
		bson.M{"$skip": (page - 1) * limit},
		bson.M{"$limit": limit},
	}
}
