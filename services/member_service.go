package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UserService service layer
type MemberService struct {
	logger     lib.Logger
	repository repository.MemberRepository
}

func NewMemberService(logger lib.Logger, repository repository.MemberRepository) domains.MemberService {
	return MemberService{
		logger:     logger,
		repository: repository,
	}
}

func (s MemberService) WithTrx(trxHandle *gorm.DB) domains.MemberService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s MemberService) ListMemberCms() (member []map[string]interface{}, err error) {
	return member, s.repository.Table("klaklik_member.member_sso").Select("*").Scan(&member).Error
}
