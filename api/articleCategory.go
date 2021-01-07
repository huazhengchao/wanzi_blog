package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/model"
)

type ArticleCategoryController struct {

}

var articleCategoryModel = model.ArticleCategory{}

// 文章分类列表
func (a * ArticleCategoryController) List(c *gin.Context) {
	list := articleCategoryModel.List()
	c.JSON(200, SuccessResponse(Success, 0, list))
}
