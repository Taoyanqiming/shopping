package dao

import "shopping/model"

// 创建订单
func CreatOrder(order model.Order) error {
	_, err := SqlConn().Exec("insert into shopping.order  value (?,?,?,?)", order)
	return err
}
