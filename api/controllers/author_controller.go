package controllers

import (
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

// UserController data type
type AuthorController struct {
	service domains.AuthorService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewAuthorController(authorservice domains.AuthorService, logger lib.Logger) AuthorController {
	return AuthorController{
		service: authorservice,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u AuthorController) GetOneAuthor(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	author, err := u.service.GetOneAuthor(id)

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": author,
	})

}
