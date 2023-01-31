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
	"strings"
)

// 创建商品
func CreateGood(ctx *gin.Context) {

	var requestGood model.Good
	err := ctx.ShouldBind(&requestGood)

	if err != nil {
		log.Printf("创建商品数据绑定失败  error : %v", err.Error())
		return
	}

	// 获取参数
	name := requestGood.Name
	info := requestGood.Info
	number := requestGood.Number //商品数量

	if !dao.VerifyGoodName(name) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "商品名称不合法,商品名称为1-20字节的中文或英文")
		return
	}

	if len(info) > 200 {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "商品简介不合法,商品简介长度为0-200字节")
		return
	}

	if number < 0 || number > 99 {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "商品数量不合法，商品数量0-99")
		return
	}
	var g model.Good
	g.Name = name
	g.Info = info
	g.Number = number
	// 保存商品
	good := dao.InsetGood(g)
	response.Success(ctx, good, "创建成功")
}

// 上传商品图片
func UploadPic(ctx *gin.Context) {
	file, err := ctx.FormFile("img")
	if err != nil {
		response.Failure(ctx, nil, "文件上传失败")
		return
	}

	// 上传文件格式不正确
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		response.Failure(ctx, nil, "上传失败,只允许png,jpg,jpeg文件")
		return
	}
	var good model.Good
	name := ctx.Param("name")
	var db sql.DB

	row := db.QueryRow("select * from shopping.good where name= ?", name).Scan(&good)
	if row != nil {
		return
	}
	if good.Number == 0 {
		response.Failure(ctx, nil, "该商品不存在")
		return
	}
	f := dao.Head2(fileExt, good.GoodId)

	// 返回值
	response.Success(ctx, gin.H{"hash": f}, "上传图片成功")
}

// 商品列表页
/*func GetGoodsList(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	var goods []model.Good
	if tokenString == "" {
		DB.Find(&goods)
	} else {
		// 解析token
		token, claims, err := dao.ParseToken(tokenString)

		// 出现错误或者token无效
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, nil, "权限不足")
			return
		}

		// 验证通过后获取claims中的userID
		userId := claims.UserId

		DB.Find(&goods, "user_id = ?", userId)
	}

	response.Success(ctx, gin.H{"goods": goods}, "查询成功")
}*/

// 商品详情页
func GetGoodDetail(ctx *gin.Context) {
	goodId := ctx.Param("goodId")

	var good model.Good

	var db sql.DB

	row := db.QueryRow("select * from shopping.good where good_id= ?", goodId).Scan(&good)
	if row != nil {
		return
	}
	if good.Number == 0 {
		response.Failure(ctx, nil, "该商品不存在")
		return
	}

	// 返回
	response.Success(ctx, gin.H{"good": good}, "查询成功")
}

// 购买商品
func BuyGood(ctx *gin.Context) {
	// 数据绑定
	var requestOrder model.Order
	err := ctx.ShouldBind(&requestOrder)

	if err != nil {
		log.Printf("购买商品数据绑定出错 error : %v", err.Error())
		return
	}

	// 获取参数
	goodId := requestOrder.GoodId
	goodCount := requestOrder.GoodCount

	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	//查询商品信息
	b, err := dao.SelectGood(goodId)
	//查询用户信息
	a, err := dao.SelectUser(userId)
	// 查询是否存在该商品
	if b.GoodId == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "不存在该商品")
		return
	}

	// 判断商品数量是否正确
	if b.Number < goodCount {
		response.Response(ctx, http.StatusBadRequest, nil, "数据错误，该商品数量不够")
		return
	}

	// 商品数量减少 保存
	b.Number -= goodCount
	j := dao.Number(b.Number, goodId)
	fmt.Println(j)

	//填写用户信息

	// 保存订单
	order := model.Order{
		UserId:      userId,
		GoodId:      goodId,
		GoodCount:   goodCount,
		GoodPrice:   float32(goodCount) * b.Price, //计算总价
		GoodName:    b.Name,
		Address:     a.Address,
		UserAccount: a.Account,
		UserPhone:   a.Phone,
	}
	//保存订单
	k := dao.CreatOrder(order)
	fmt.Println(k)

	//余额减少
	a.Balance -= float32(goodCount) * b.Price
	l := dao.ChargeMoney(a.Account, a.Balance)
	fmt.Println(l)

	//销量+
	b.Sale += goodCount
	m := dao.Sale(b.Sale, goodId)
	fmt.Println(m)

	response.Success(ctx, gin.H{"order": order}, "购买成功")
}

// 价格降序展示商品
func PriceDESC(ctx *gin.Context) {
	name := ctx.Param("name")
	var good model.Good
	var db sql.DB
	row := db.QueryRow("select * from shopping.good where name= ? order by price DESC ", name).Scan(&good)
	if row != nil {
		return
	}
	if good.Number == 0 {
		response.Failure(ctx, nil, "该商品不存在")
		return
	}

	// 返回
	response.Success(ctx, gin.H{"good": good}, "查询成功")

}

// 价格升序
func PriceASC(ctx *gin.Context) {
	name := ctx.Param("name")
	var good model.Good
	var db sql.DB
	row := db.QueryRow("select * from shopping.good where name= ? order by price  ", name).Scan(&good)
	if row != nil {
		return
	}
	if good.Number == 0 {
		response.Failure(ctx, nil, "该商品不存在")
		return
	}

	// 返回
	response.Success(ctx, gin.H{"good": good}, "查询成功")

}
