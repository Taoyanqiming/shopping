package dao

import "shopping/model"

// 创建订单
func CreatOrder(order model.Order) error {
	_, err := SqlConn().Exec("insert into shopping.order  value (?,?,?,?,?,?,?,?,?)", order)
	return err
}

// 删除
func CancelOrder(userId int) error {
	_, err := SqlConn().Exec("delete from shopping.cart where user_id=? ", userId)
	return err
}
