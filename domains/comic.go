package domains

import (
	"github.com/dipeshdulal/clean-gin/models"
	"gorm.io/gorm"
)

type ComicService interface {
	WithTrx(trxHandle *gorm.DB) ComicService
	GetListCms() ([]models.Comic, error)
}
