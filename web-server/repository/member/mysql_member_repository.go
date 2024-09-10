package member

import (
	"fmt"
	"web-server/models/member"

	"gorm.io/gorm"
)

type MysqlMemberRepository struct {
	DB *gorm.DB
}

func (r *MysqlMemberRepository) Save(member *member.Member) error {
	fmt.Println(member)

	return r.DB.Create(member).Error
}

func (r *MysqlMemberRepository) FindByID(id uint)(*member.Member, error) {
	var member member.Member
	err := r.DB.First(&member, id).Error

	if(err != nil){
		return nil, err
	}

	return &member, nil
}