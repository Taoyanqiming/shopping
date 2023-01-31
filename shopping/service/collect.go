package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
)

// 收藏
func Collect(ctx *gin.Context) {
	var Collect model.Collect
	err := ctx.ShouldBind(&Collect)

	if err != nil {
		log.Printf("绑定失败  error : %v", err.Error())
		return
	}
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	com := model.Collect{
		UserId: userId,
		GoodId: Collect.GoodId,
	}

	c := dao.Collect(com)
	fmt.Println(c)
	response.Success(ctx, c, "收藏成功")
}

// 查看
// 判断是否已经收藏
// 删除
func DeleteCol(ctx *gin.Context) {
	var Collect model.Collect
	err := ctx.ShouldBind(&Collect)

	if err != nil {
		log.Printf("绑定失败  error : %v", err.Error())
		return
	}

	goodId := Collect.GoodId

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	d := dao.Cancel(userId, goodId)
	fmt.Println(d)

	response.Success(ctx, nil, "收藏删除成功")
}
