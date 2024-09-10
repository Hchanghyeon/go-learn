package member

import (
	"web-server/models/member"
)

type MemberRepository interface {
	Save(member *member.Member) error
	FindByID(id uint) (*member.Member, error)	
}