package model

type ArticleCategory struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// 文章分类列表
func (a *ArticleCategory) List() []ArticleCategory {
	var articleCategory []ArticleCategory
	DB.Find(&articleCategory)
	return articleCategory
}