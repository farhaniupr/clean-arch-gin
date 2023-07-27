package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type MemberRoutes struct {
	logger           lib.Logger
	handler          lib.RequestHandler
	memberController controllers.MemberController
	authMiddleware   middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s MemberRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/list-member", s.memberController.ListMember)
	}
}

// NewUserRoutes creates new user controller
func NewMemberRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	memberController controllers.MemberController,
	authMiddleware middlewares.JWTAuthMiddleware,
) MemberRoutes {
	return MemberRoutes{
		handler:          handler,
		logger:           logger,
		memberController: memberController,
		authMiddleware:   authMiddleware,
	}
}
