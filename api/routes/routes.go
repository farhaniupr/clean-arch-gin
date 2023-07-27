package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewAuthorRoutes),
	fx.Provide(NewComicRoutes),
	fx.Provide(NewMemberRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	authRoutes AuthRoutes,
	authorRoutes AuthorRoutes,
	comicRoutes ComicRoutes,
	memberRoutes MemberRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		authorRoutes,
		comicRoutes,
		memberRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
