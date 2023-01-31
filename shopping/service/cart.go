package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
)

// 添加购物车
func CreateCart(ctx *gin.Context) {

	var requestCart model.Cart
	err := ctx.ShouldBind(&requestCart)

	if err != nil {
		log.Printf("创建商品数据绑定失败  error : %v", err.Error())
		return
	}

	// 获取参数(do not know where to get)
	goodId := requestCart.GoodId

	if !dao.IsCartExist(goodId) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "商品不存在")
		return
	}
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	//查询商品信息
	b, err := dao.SelectGood(goodId)
	//查询用户信息
	a, err := dao.SelectUser(userId)
	var l = model.Cart{
		Userid:      userId,
		GoodId:      goodId,
		UserAccount: a.Account,
		GoodName:    b.Name,
		GoodPrice:   b.Price,
		GoodNumber:  b.Number,
		GoodInfo:    b.Info,
	}
	// 加入购物车
	z := dao.InsertCart(l)
	fmt.Println(z)
	response.Success(ctx, gin.H{"cart": l}, "创建成功")
}

// 部分商品删除
// 清空购物车
func EmptyCart(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	d := dao.DeleteCart(userId)
	fmt.Println(d)
	response.Success(ctx, nil, "购物车清空")
}
