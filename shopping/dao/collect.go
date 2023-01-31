package dao

import (
	"shopping/model"
)

// 收藏商品
func Collect(collect model.Collect) error {
	_, err := SqlConn().Exec("insert into shopping.collect  value (?,?,?)", collect)
	return err
}

// 取消收藏
func Cancel(userId, goodId int) error {
	_, err := SqlConn().Exec("delete from shopping.collect where use_id=? and good_id = ?", userId, goodId)
	return err
}

// 查询收藏
func SelectCollect(userId, goodId int) (model.Collect, error) {
	var collect = model.Collect{}
	row := db.QueryRow("select * from shopping.collect where good_id=? and user_id= ?", goodId, userId)
	if row != nil {
		return collect, row.Err()
	}
	err := row.Scan(&collect)
	if err != nil {
		return collect, err
	}
	return collect, nil
}

// 检查是否已经收藏
/*func IsCollectExist(userId,goodId int) bool {
	c, err := SelectCollect(userId,goodId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	if c.GoodId != goodId {
		return false
	}
	return true
}*/
