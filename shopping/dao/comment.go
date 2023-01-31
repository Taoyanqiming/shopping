package dao

import (
	"shopping/model"
)

// 评论
func Comment(comment model.Comment) error {
	_, err := SqlConn().Exec("insert into shopping.comment  value (?,?,?)", comment.UserId, comment.GoodId, comment.Comment)
	return err
}

// 删
func DeleteCom(userId, goodId int) error {
	_, err := SqlConn().Exec("delete from shopping.comment  where user_id=?and good_id =?", userId, goodId)
	return err
}
