package dao

import (
	"database/sql"
	"shopping/model"
	"unicode"
)

// 验证商品名是否合法，商品名只能由1-20位中文或英文构成
func VerifyGoodName(name string) bool {
	if len(name) < 1 || len(name) > 20 {
		return false
	} else {
		flag := true
		// 店铺名只能由中文和英文构成
		for _, r := range name {
			if !unicode.IsLetter(r) {
				flag = false
			}
		}
		return flag
	}
}

// 查询商品
func SelectGood(goodId int) (model.Good, error) {
	good := model.Good{}
	row := db.QueryRow("select * from shopping.good where good_id= ?", goodId)
	if row != nil {
		return good, row.Err()
	}
	err := row.Scan(&good)
	if err != nil {
		return good, err
	}
	return good, nil
}

// 创建商品
func InsetGood(good model.Good) error {
	_, err := SqlConn().Exec("insert into shopping.good (name,number,price,info) value (?,?,?,?)", good.Name, good.Number, good.Price, good.Info)
	return err
}

// 商品存在？
func IsGoodExist(goodId int) bool {
	good, err := SelectGood(goodId)
	if err != nil {
		if err == sql.ErrNoRows {
			//查询是否返回数据为空，用于？？
			return false
		}
		return false
	}
	if good.GoodId == goodId {

		return false
	}
	return true
}

// 更新商品数量
func Number(number, goodId int) error {
	_, err := SqlConn().Exec("update shopping.good set number =? where good_id=?", number, goodId)
	return err
}

// 销量更新
func Sale(sale, goodId int) error {
	_, err := SqlConn().Exec("update shopping.good set sale =? where good_id=?", sale, goodId)
	return err
}

// 图片
func Head2(head string, goodId int) error {
	_, err := SqlConn().Exec("update shopping.good set picture=? where good_id=?", head, goodId)
	return err
}
