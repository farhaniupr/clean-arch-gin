package controllers

import (
	"net/http"

	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

// UserController data type
type MemberController struct {
	service domains.MemberService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewMemberController(memberservice domains.MemberService, logger lib.Logger) MemberController {
	return MemberController{
		service: memberservice,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u MemberController) ListMember(c *gin.Context) {

	member, err := u.service.ListMemberCms()

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": member,
	})

}
