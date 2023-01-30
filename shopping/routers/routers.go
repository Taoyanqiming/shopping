package routers

import (
	"github.com/gin-gonic/gin"
	"shopping/middleware"
	"shopping/service"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 用户
	r.POST("/user/register", service.Register)
	r.POST("/user/login", service.Login)
	r.PUT("/user/change/info", middleware.AuthMiddleware(), service.UpdateUser)         //改姓名，头像
	r.POST("/user/upload/head", middleware.AuthMiddleware(), service.UploadHead)        //上传头像
	r.GET("/user/info", middleware.AuthMiddleware(), service.GetUserData)               //获取用户信息
	r.PUT("/user/change/password", middleware.AuthMiddleware(), service.UpdatePassword) //改密码
	r.GET("/user/check/balance", middleware.AuthMiddleware(), service.Check)            //余额查询
	r.PUT("/user/charge/money", middleware.AuthMiddleware(), service.Charge)            //充值
	//商品收藏

	//商品
	r.POST("/api/good", middleware.AuthMiddleware(), service.CreateGood)       //创建商品
	r.GET("/api/goods/:goodId", service.GetGoodDetail)                         //商品详情页
	r.POST("/api/upload/good", middleware.AuthMiddleware(), service.UploadPic) //上传商品图片
	//商品评论
	//商品排序

	//购物车
	//添加购物车
	//删除商品
	//情况购物车

	//订单
	r.PUT("/order/address", middleware.AuthMiddleware(), service.Address) //选择收获地址
	r.PUT("/order", middleware.AuthMiddleware(), service.BuyGood)         //提交订单
	//取消订单
	//确认收获
	//商品评价
	return r
}
