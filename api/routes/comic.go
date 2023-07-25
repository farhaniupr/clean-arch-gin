package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type ComicRoutes struct {
	logger          lib.Logger
	handler         lib.RequestHandler
	comicController controllers.ComicController
	authMiddleware  middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s ComicRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/list-comic", s.comicController.GetListCms)
	}
}

// NewUserRoutes creates new user controller
func NewComicRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	comiccontroller controllers.ComicController,
	authMiddleware middlewares.JWTAuthMiddleware,
) ComicRoutes {
	return ComicRoutes{
		handler:         handler,
		logger:          logger,
		comicController: comiccontroller,
		authMiddleware:  authMiddleware,
	}
}
