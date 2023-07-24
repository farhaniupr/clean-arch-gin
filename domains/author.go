package domains

import (
	"github.com/dipeshdulal/clean-gin/models"
	"gorm.io/gorm"
)

type AuthorService interface {
	WithTrx(trxHandle *gorm.DB) AuthorService
	GetOneAuthor(id int) (models.Author, error)
}
