package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// UserRepository database structure
type AuthorRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func AuthorConnectRepository(db lib.Database, logger lib.Logger) AuthorRepository {
	return AuthorRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r AuthorRepository) WithTrx(trxHandle *gorm.DB) AuthorRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
