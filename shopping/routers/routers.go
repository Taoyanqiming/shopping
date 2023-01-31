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
	r.PUT("/user/change/info", middleware.AuthMiddleware(), service.UpdateUser)         //改姓名，性别
	r.POST("/user/upload/head", middleware.AuthMiddleware(), service.UploadHead)        //上传头像
	r.GET("/user/info", middleware.AuthMiddleware(), service.GetUserData)               //获取用户信息
	r.PUT("/user/change/password", middleware.AuthMiddleware(), service.UpdatePassword) //改密码
	r.GET("/user/check/balance", middleware.AuthMiddleware(), service.Check)            //余额查询
	r.PUT("/user/charge/money", middleware.AuthMiddleware(), service.Charge)            //充值

	//商品收藏
	r.POST("/good/collect", middleware.AuthMiddleware(), service.Collect)            //收藏
	r.DELETE("/good/delete/collect", middleware.AuthMiddleware(), service.DeleteCol) //取消收藏

	//商品
	r.POST("/good", service.CreateGood)          //创建商品
	r.GET("/good/goodId", service.GetGoodDetail) //商品详情页
	r.POST("/upload/good", service.UploadPic)    //上传商品图片

	//商品评论
	r.POST("/good/comment", middleware.AuthMiddleware(), service.Comment)            //评论
	r.DELETE("/good/delete/comment", middleware.AuthMiddleware(), service.DeleteCom) //取消评论

	//商品排序展示
	r.GET("/good/priceDESC", service.PriceDESC) //价格降序展示
	r.GET("/good/priceASC", service.PriceASC)   //价格升序展示

	//购物车
	r.POST("/cart", middleware.AuthMiddleware(), service.CreateCart) //添加购物车
	//删除部分商品
	r.DELETE("/delete/cart", middleware.AuthMiddleware(), service.EmptyCart) //清空购物车

	//订单
	r.PUT("/order/address", middleware.AuthMiddleware(), service.Address)       //选择收获地址
	r.PUT("/order", middleware.AuthMiddleware(), service.BuyGood)               //提交订单
	r.DELETE("/delete/order", middleware.AuthMiddleware(), service.CancelOrder) //取消订单
	//确认收获
	return r
}
