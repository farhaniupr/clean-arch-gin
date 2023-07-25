package controllers

import (
	"net/http"

	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

type ComicController struct {
	service domains.ComicService
	logger  lib.Logger
}

func NewComicController(comicservice domains.ComicService, logger lib.Logger) ComicController {
	return ComicController{
		service: comicservice,
		logger:  logger,
	}
}

func (u ComicController) GetListCms(c *gin.Context) {

	comic, err := u.service.GetListCms()

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": comic,
	})

}
