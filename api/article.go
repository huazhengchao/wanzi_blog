package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/model"
	"strconv"
)

type ArticleController struct {
}

var articlesModel model.Article

// 文章列表
func (a *ArticleController) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultPostForm("page", "1"))
	title := c.DefaultPostForm("title", "")
	if err != nil {
		c.JSON(200, ErrorResponse(ParamError))
	}
	list := articlesModel.List(page, title)
	c.JSON(200, SuccessResponse(Success, list))
}

// 文章详情
func (a *ArticleController) Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, ErrorResponse(ParamError))
	}
	data := articlesModel.Detail(id)
	c.JSON(200, SuccessResponse(Success, 0, data))
}