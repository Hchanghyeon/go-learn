package member

type Member struct {
	Id uint `json:"id"`
	UserId string `json:"userId"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Age uint8 `json:"age"`
}

type CreateMemberRequest struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Age uint8 `json:"age"`
}