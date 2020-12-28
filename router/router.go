package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/api"
)

func NewRouter() *gin.Engine{

	router := gin.Default()

	articleController := new(api.ArticleController)
	// 文章列表接口
	router.POST("/getList", articleController.List)

	return router
}
