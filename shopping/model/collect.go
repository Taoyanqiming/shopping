package model

type Collect struct {
	UserId      int    `json:"user_id"`
	GoodId      int    `json:"good_id"`
	GoodName    string `json:"good_name"`
	UserAccount int    `json:"user_account"`
}
