package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
)

// 评论
func Comment(ctx *gin.Context) {
	var Comment model.Comment
	err := ctx.ShouldBind(&Comment)

	if err != nil {
		log.Printf("绑定失败  error : %v", err.Error())
		return
	}

	com := model.Comment{
		UserId:  Comment.UserId,
		GoodId:  Comment.GoodId,
		Comment: Comment.Comment,
	}

	c := dao.Comment(com)
	fmt.Println(c)
	response.Success(ctx, c, "评论成功")
}

// 查看
// 删除
func DeleteCom(ctx *gin.Context) {
	var Comment model.Comment
	err := ctx.ShouldBind(&Comment)
	if err != nil {
		log.Printf("绑定失败  error : %v", err.Error())
		return
	}

	goodId := Comment.GoodId
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId

	d := dao.DeleteCom(userId, goodId)
	fmt.Println(d)

	response.Success(ctx, nil, "评论删除成功")
}
