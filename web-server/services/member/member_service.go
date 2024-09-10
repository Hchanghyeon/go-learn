package member

import (
	"fmt"
	dbConfig "web-server/config"
	"web-server/models/member"
	repository "web-server/repository/member"
)

var memberRepository repository.MemberRepository

func init(){
	db := dbConfig.GetDB()

	if(db == nil){
		fmt.Println("error");
	}

	memberRepository = &repository.MysqlMemberRepository{DB: db}
}

func CreateMember(createMemberRequest *member.CreateMemberRequest)(member.Member){
	
	member := member.Member{
		Nickname: createMemberRequest.Nickname,
		Password:createMemberRequest.Password,
		UserId: createMemberRequest.UserId,
		Age: createMemberRequest.Age,
	}

	memberRepository.Save(&member)

	return member
} 