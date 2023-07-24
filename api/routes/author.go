package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type AuthorRoutes struct {
	logger           lib.Logger
	handler          lib.RequestHandler
	authorController controllers.AuthorController
	authMiddleware   middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s AuthorRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/author/:id", s.authorController.GetOneAuthor)
	}
}

// NewUserRoutes creates new user controller
func NewAuthorRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	authorController controllers.AuthorController,
	authMiddleware middlewares.JWTAuthMiddleware,
) AuthorRoutes {
	return AuthorRoutes{
		handler:          handler,
		logger:           logger,
		authorController: authorController,
		authMiddleware:   authMiddleware,
	}
}
