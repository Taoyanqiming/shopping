package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
)

func CancelOrder(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	c := dao.CancelOrder(userId)
	fmt.Println(c)
	response.Success(ctx, nil, "订单已取消")

}
