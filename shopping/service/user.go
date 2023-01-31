package service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
	"shopping/utils"
	"strings"
)

// 注册
func Register(ctx *gin.Context) {

	var requestUser = model.User{}

	// 绑定数据
	err := ctx.ShouldBind(&requestUser)

	if err != nil {
		log.Printf("注册数据绑定失败 error : %v", err.Error())
		return
	}
	// 获取数据
	account := requestUser.Account
	password := requestUser.Password
	name := utils.RandomString(5) //随机生成昵称

	// 数据验证
	if !dao.VerifyAccount(account) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "账号不合法，账号长度为6-20字符，可以使用英文和数字")
		return
	}

	if !dao.VerifyPassword(password) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "密码不合法，账号长度为8-16字符，可以使用英文、数字和特殊字符")
		return
	}

	// 验证账号是否已经存在
	/*if !dao.IsAccountExist(account) {
		response.Response(ctx, http.StatusBadRequest, nil, "该账号已经存在")
		return
	}*/

	// 保存新用户到数据库
	newUser := model.User{
		Account:  account,
		Password: password,
	}
	// 保存新用户到数据库
	user := dao.Register(account, password, name)
	response.Success(ctx, user, "账号密码已保存")

	// 发放token
	token, err := dao.IssueToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, nil, "服务端出错")
		log.Printf("注册时生成token失败 error : %v", err.Error())
		return
	}

	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

// 登录
func Login(ctx *gin.Context) {
	var requestUser = model.User{}

	// 绑定对象，参数为指针
	err := ctx.ShouldBind(&requestUser)
	if err != nil {
		log.Printf("登录绑定对象失败 error : %v", err.Error())
		return
	}

	// 获取参数
	account := requestUser.Account
	password := requestUser.Password

	// 数据验证
	if !dao.VerifyAccount(account) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "账号不合法，账号长度为6-20字符，可以使用英文和数字")
		return
	}

	if !dao.VerifyPassword(password) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "密码长度为8-16字符，可以使用英文、数字和特殊字符")
		return
	}

	// 判断账号是否已经存在
	if dao.IsAccountExist(account) {
		response.Response(ctx, http.StatusBadRequest, nil, "该账号不存在")
		return
	}

	// 判断密码是否正确
	if !dao.IsPasswordCorrect(account, password) {
		response.Failure(ctx, nil, "密码不正确")
		return
	}

	var user model.User

	// 生成token
	token, err := dao.IssueToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, nil, "系统异常,生成token失败")
		log.Printf("token 获取异常 error : %v", err.Error())
		return
	}
	// 登录成功
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

// 修改用户信息
func UpdateUser(ctx *gin.Context) {
	var requestUser model.User

	// 数据绑定
	err := ctx.ShouldBind(&requestUser)

	if err != nil {
		log.Printf("更新用户绑定出错 error : %v", err.Error())
		return
	}

	// 获取数据
	name := requestUser.Name
	sex := requestUser.Sex

	// 昵称长度 1-20
	if len(name) < 1 || len(name) > 20 {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "昵称格式不正确，昵称长度为1-20字符")
		return
	}

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	// 修改该用户的个人信息
	var User model.User
	User.Name = name
	User.Sex = sex

	// 保存记录
	user1 := dao.ChangeUserInfo(name, sex, userId)

	// 更新用户信息成功
	response.Success(ctx, gin.H{"user": user1}, "更新用户信息成功")
}

// 获取用户信息
func GetUserData(ctx *gin.Context) {
	user, isExist := ctx.Get("user")

	if isExist {
		response.Success(ctx, gin.H{"user": user}, "获取成功")
	} else {
		return
	}
}

// 上传头像
func UploadHead(ctx *gin.Context) {
	f, err := ctx.FormFile("img")

	// 上传失败
	if err != nil {
		response.Failure(ctx, nil, "上传失败")
		return
	}

	// 上传文件格式不正确
	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		response.Failure(ctx, nil, "上传失败,只允许png,jpg,jpeg文件")
		return
	}

	// 保存成功，修改用户头像字段
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	// 更新到数据库
	var updateUser model.User

	updateUser.Head = fileExt
	h := dao.Head(updateUser.Head, userId)

	// 返回值
	response.Success(ctx, h, "上传图片成功")
}

// 改密码
func UpdatePassword(ctx *gin.Context) {
	type Password struct {
		new string
		old string
	}
	var requestUser Password

	// 数据绑定
	err := ctx.ShouldBind(&requestUser)

	if err != nil {
		log.Printf("更新用户绑定出错 error : %v", err.Error())
		return
	}

	// 获取数据
	oldPassword := requestUser.old
	newPassword := requestUser.new
	//token
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	// 将该用户取出
	var db *sql.DB
	var User model.User
	row := db.QueryRow("select * from shopping.user where user_id= ?", userId)
	err = row.Scan(&User.Account)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(User)

	// 判断密码是否正确（费事了）
	if !dao.IsPasswordCorrect(User.Account, oldPassword) {
		response.Failure(ctx, nil, "密码不正确")
		return
	}

	// 判断新密码是否符合条件
	if !dao.VerifyPassword(newPassword) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "密码不合法，账号长度为8-16字符，可以使用英文、数字和特殊字符")
		return
	}

	h := dao.ChangePassword(oldPassword, newPassword, userId)
	fmt.Println(h)
	// 更新用户信息成功
	response.Success(ctx, gin.H{}, "更新用户信息成功")
}

// 氪金
func Charge(ctx *gin.Context) {
	var requestUser model.User

	// 数据绑定
	err := ctx.ShouldBind(&requestUser)

	if err != nil {
		log.Printf("更新用户绑定出错 error : %v", err.Error())
		return
	}
	//获取数据
	balance := requestUser.Balance

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	var db *sql.DB
	var User model.User
	row := db.QueryRow("select * from shopping.user where user_id= ?", userId)
	err1 := row.Scan(&User.Balance, &User.Account)
	if err1 != nil {
		log.Println(err1)
		return
	}
	fmt.Println(User)

	var newUser model.User
	newUser.Balance = User.Balance + balance
	//保存数据
	i := dao.ChargeMoney(User.Account, newUser.Balance)
	fmt.Println(i)
	response.Success(ctx, gin.H{}, "用户充值成功")
}

// 查询余额
func Check(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	var db *sql.DB
	row := db.QueryRow("select balance from shopping.user where user_id= ?", userId)
	err := row.Scan(&user)
	if err != nil {
		fmt.Println("")
	}
	fmt.Println("")
	response.Success(ctx, err, "查询成功")
}

// 填写地址
func Address(ctx *gin.Context) {
	var requestUser = model.User{}
	// 绑定数据
	err := ctx.ShouldBind(&requestUser)

	if err != nil {
		log.Printf("注册数据绑定失败 error : %v", err.Error())
		return
	}

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	// 获取数据
	address := requestUser.Address
	phone := requestUser.Phone
	a := dao.Address(address, phone, userId)
	response.Success(ctx, gin.H{"a": a}, "填写成功")
}
