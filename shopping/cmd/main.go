package main

import (
	"github.com/gin-gonic/gin"
	"shopping/dao"
	"shopping/routers"
)

func main() {
	dao.SqlConn()
	r := gin.Default()
	r = routers.CollectRoute(r)
	r.Run()
}
