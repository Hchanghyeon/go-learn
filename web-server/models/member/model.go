package member

type Member struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
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