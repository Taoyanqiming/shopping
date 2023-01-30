package model

// 订单
type Order struct {
	UserId      int     `json:"user_id"`
	GoodId      int     `json:"good_id"`
	Address     string  `json:"address"`
	GoodName    string  `json:"good_name"`
	GoodPrice   float32 `json:"good_price"`
	GoodCount   int     `json:"good_count"` //购买数量
	UserAccount string  `json:"user_account"`
	UserPhone   int     `json:"user_phone"`
}
