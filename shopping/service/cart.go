package service

// 添加购物车
/*func CreateCart(ctx *gin.Context) {

	var requestCart model.Cart
	err := ctx.ShouldBind(&requestCart)

	if err != nil {
		log.Printf("创建商品数据绑定失败  error : %v", err.Error())
		return
	}

	// 获取参数(do not know where to come)
	goodId := requestCart.GoodId

	if dao.IsCartExist(goodId) {
		response.Response(ctx, http.StatusUnprocessableEntity, nil, "商品存在")
		return
	}
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	account := user.(model.User).Account
	var a = model.Cart{
		Userid:      userId,
		GoodId:      goodId,
		UserAccount: account,
		GoodName:    "",
		GoodPrice:   0,
		GoodNumber:  0,
		GoodInfo:    "",
	}
	dao.
		// 保存商品

		cart1 = dao.InsertCart(a)
	response.Success(ctx, gin.H{"cart": cart1}, "创建成功")
}*/
//清空购物车
//部分商品删除
