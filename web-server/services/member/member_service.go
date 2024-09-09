package member

import (
	"web-server/models/member"
)

func CreateMember(createMemberRequest member.CreateMemberRequest)(member.Member){
	member := member.Member{
		Id: 1,
		UserId: createMemberRequest.UserId,
		Password: createMemberRequest.Password,
		Nickname: createMemberRequest.Nickname,
		Age: createMemberRequest.Age,
	}

	return member
}