package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// UserRepository database structure
type MemberRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewMemberRepository(db lib.Database, logger lib.Logger) MemberRepository {
	return MemberRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r MemberRepository) WithTrx(trxHandle *gorm.DB) MemberRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
