package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/api"
	"github.com/huazhengchao/wanzi_blog/middleware"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	// 解决跨域
	router.Use(middleware.Cors())

	articleController := new(api.ArticleController)
	// 文章列表接口
	router.POST("/getList", articleController.List)
	// 文章详情接口
	router.POST("/detail", articleController.Detail)

	articleMessageController := new(api.ArticleMessageController)
	// 文章留言列表接口
	router.POST("/messageList", articleMessageController.List)
	// 文章留言新增接口
	router.POST("/messageAdd", articleMessageController.Add)

	articleCategoryController := new(api.ArticleCategoryController)
	// 文章分类列表
	router.GET("/categoryList", articleCategoryController.List)

	return router
}
