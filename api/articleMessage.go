package api

import (
	"github.com/gin-gonic/gin"
	"github.com/huazhengchao/wanzi_blog/model"
	"strconv"
)

type ArticleMessageController struct {
}

var articleMessageModel model.ArticleMessage

// 文章留言列表
func (a *ArticleMessageController) List(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(200, ErrorResponse(ParamError))
	}
	list := articleMessageModel.List(id)
	c.JSON(200, SuccessResponse(Success, 0, list))
}
