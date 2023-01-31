package dao

import (
	"fmt"
	"shopping/model"
	"unicode"
)

func VerifyAccount(account string) bool {
	// 若长度不符合直接返回false
	if len(account) < 6 || len(account) > 20 {
		return false
	} else {
		flag := true
		// 判断每一个字符 不是英文并且也不是数字 则不合法
		for _, r := range account {
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
				flag = false
			}
		}
		return flag
	}
}

// 验证密码是否合法 长度 8-16字符 只能英文 数字 特殊字符
func VerifyPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {
		return false
	} else {
		flag := true
		// 判断密码是否由 数字 英文 特殊字符
		for _, r := range password {
			if !unicode.IsLetter(r) && !unicode.IsNumber(r) && (r < 33 || r > 47) {
				flag = false
			}
		}
		return flag
	}
}

// 插入数据，用于注册用户
func Register(account, password, name string) error {
	_, err := SqlConn().Exec("insert into shopping.user (account ,password,name) value (?,?,?)", account, password, name)
	return err
}

// 根据userid查询用户
func SelectUser(userid int) (model.User, error) {
	user := model.User{}
	row := db.QueryRow("select* from shopping.user where user_id= ?", userid)
	if row != nil {
		return user, row.Err()
	}
	err := row.Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func SelectUser2(account string) (model.User, error) {
	user := model.User{}
	row := db.QueryRow("select * from shopping.user where account= ?", account)
	if row != nil {
		return user, row.Err()
	}
	err := row.Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// 用户存在？
func IsAccountExist(account string) bool {
	user, err := SelectUser2(account)
	if err != nil {
		fmt.Println("账号不存在")
	}
	if user.UserId != 0 {
		return true
	}
	return false

}

// 验证密码是否正确
func IsPasswordCorrect(account, password string) bool {
	user, err := SelectUser2(account)
	if err != nil {
		fmt.Println("账号不存在")
	}
	if user.Password != password {
		return false
	}
	return true
}

// 数据库改变用户基本信息
func ChangeUserInfo(name string, sex byte, userid int) error {
	_, err := SqlConn().Exec("update shopping.user set name=? ,sex=?where user_id=?", name, sex, userid)
	return err
}

// 改密码
func ChangePassword(old, new string, userid int) error {
	_, err := db.Exec("update shopping.user set password=? where user_id=?", new, old, userid)
	return err
}

func Head(head string, userid int) error {
	_, err := SqlConn().Exec("update shopping.user set head=? where user_id=?", head, userid)
	return err
}

// 余额查询
func SelectBalance(userid int) (model.User, error) {
	user := model.User{}
	row := db.QueryRow("select balance from shopping.user where user_id= ?", userid)
	if row != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Balance)
	if err != nil {
		return user, err
	}
	return user, nil
}

// 钱包变动
func ChargeMoney(account string, balance float32) error {
	_, err := SqlConn().Exec("update shopping.user set balance=? where account=?", balance, account)
	return err
}

// 输入地址
func Address(address string, phone int, userId int) error {
	_, err := SqlConn().Exec("update shopping.user set address=?,phone =? where account=?", address, phone, userId) //是否为指定账户,如上
	return err
}
