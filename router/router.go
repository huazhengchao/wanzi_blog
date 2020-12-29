package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/api"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	articleController := new(api.ArticleController)
	// 文章列表接口
	router.POST("/getList", articleController.List)
	// 文章详情接口
	router.POST("/detail", articleController.Detail)

	articleMessageController := new(api.ArticleMessageController)
	// 文章留言接口
	router.POST("/messageList", articleMessageController.List)
	return router
}
