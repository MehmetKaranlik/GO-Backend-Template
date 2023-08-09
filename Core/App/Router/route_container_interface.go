package Router

type IRouteContainer interface {
	MakeRoutes() []Route
}
