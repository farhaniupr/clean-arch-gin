package domains

import (
	"gorm.io/gorm"
)

type MemberService interface {
	WithTrx(trxHandle *gorm.DB) MemberService
	ListMemberCms() ([]map[string]interface{}, error)
}
