package model

type Cart struct {
	Userid      int    `json:"user_id"`
	GoodId      int    `json:"good_id"`
	UserAccount string `json:"user_account"`
	GoodName    string `json:"good_name"`
	GoodPrice   int    `json:"good_price"`
	GoodNumber  int    `json:"good_number"`
	GoodInfo    string `json:"good_info"`
}
