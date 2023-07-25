package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

type ComicService struct {
	logger     lib.Logger
	repository repository.ComicRepository
}

func NewComicService(logger lib.Logger, repository repository.ComicRepository) domains.ComicService {
	return ComicService{
		logger:     logger,
		repository: repository,
	}
}

func (s ComicService) WithTrx(trxHandle *gorm.DB) domains.ComicService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s ComicService) GetListCms() (comic []models.Comic, err error) {
	return comic, s.repository.Table("comic").Select("*").Scan(&comic).Error
}
