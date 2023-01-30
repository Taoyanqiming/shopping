package model

type Good struct {
	GoodId  int     `json:"good-id" `
	Name    string  `json:"name"`
	Number  int     `json:"number" ` // 商品库存数量
	Price   float32 `json:"price" `
	Info    string  `json:"info" `
	Picture string  `json:"picture" `
	Sale    int     `json:"sale"`    //销量
	Comment int     `json:"comment"` //评价
}
