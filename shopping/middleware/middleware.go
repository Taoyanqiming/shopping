package middleware

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shopping/dao"
	"shopping/model"
	"shopping/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取Header中Authorization
		tokenString := ctx.GetHeader("Authorization")

		// 验证格式
		if tokenString == "" {
			response.Response(ctx, http.StatusUnauthorized, nil, "权限不足1")
			return
		}

		// 解析token
		token, claims, err := dao.ParseToken(tokenString)

		// 出现错误或者token无效
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, nil, "权限不足2")
			return
		}

		// 验证通过后获取claims中的userID
		userId := claims.UserId
		// 将该用户记录查出来
		var db *sql.DB
		var user model.User

		row := db.QueryRow("select * from shopping.user where user.id= ?", userId)
		err = row.Scan(&user)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(user)

		// 验证用户是否存在
		if user.UserId == 0 {
			response.Response(ctx, http.StatusUnauthorized, nil, "权限不足3")
			return
		}
		// 用户存在，则将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
