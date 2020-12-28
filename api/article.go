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
	if err != nil {
		c.JSON(200, ErrorResponse(ParamError))
	}
	list, total := articlesModel.List(page)
	c.JSON(200, SuccessResponse(Success, int(total), list))
}