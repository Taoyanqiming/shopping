package model

type User struct {
	UserId   int     `json:"user_id"  form:"user_id"`
	Account  string  `json:"account"  form:"account" `
	Password string  `json:"password" form:"password"`
	Name     string  `json:"name"     form:"name"`
	Sex      byte    `json:"sex"      form:"sex"`
	Head     string  `json:"head"  `
	Balance  float32 `json:"balance"` //余额
	Address  string  `json:"address"`
	Phone    int     `json:"phone"`
}
