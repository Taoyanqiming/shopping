package dao

import (
	"database/sql"
	"shopping/model"
)

// 查询购物车内容，可用于结算
func SelectCart(db *sql.DB, userId int) (model.Cart, error) {
	cart := model.Cart{}
	row := db.QueryRow("select * from shopping.cart where user_id= ?", userId)
	if row != nil {
		return cart, row.Err()
	}
	err := row.Scan(&cart)
	if err != nil { //购物车为空？？
		return cart, err
	}
	return cart, nil
}

// 检验商品(与订单处连接，查询购物车内商品是否存在，传入商品名称）
func IsCartExist(goodId int) bool {
	cart, err := SelectCart(db, goodId)
	if err != nil {
		if err == sql.ErrNoRows {
			//查询是否返回数据为空，用于？？
			return false
		}
		return false
	}
	if cart.GoodId == goodId {

		return false
	}
	return true

}

// 商品加入购物车
func InsertCart(cart model.Cart) error {
	_, err := SqlConn().Exec("insert into shopping.cart  value (?,?,?,?,?,?)", cart.Userid, cart.GoodId, cart.GoodName, cart.UserAccount, cart.GoodNumber, cart.GoodPrice)
	return err
}

// 清空购物车
func DeleteCart(cart model.Cart) error {
	_, err := SqlConn().Exec("delete from shopping.cart where user_id=?", cart.Userid)
	return err
}

// - 删除商品
func DeleteCart01(cart model.Cart) error {
	_, err := SqlConn().Exec("delete from shopping.cart where good_id=? and user_id", cart.GoodId, cart.Userid)
	return err
}

// - 选择部分商品进行结算(传入订单数据库）
func Pay(good_id int) {

}
