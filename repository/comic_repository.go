package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// UserRepository database structure
type ComicRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewComicRepository(db lib.Database, logger lib.Logger) ComicRepository {
	return ComicRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (c ComicRepository) WithTrx(trxHandle *gorm.DB) ComicRepository {
	if trxHandle == nil {
		c.logger.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.Database.DB = trxHandle
	return c
}
