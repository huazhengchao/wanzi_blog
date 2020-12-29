package model

import (
	"time"
)

type ArticleMessage struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// 文章列表
func (a *ArticleMessage) List(id int) []ArticleMessage {
	var articles []ArticleMessage
	DB.Where("pid = ?", id).
		Order("created_at desc").
		Find(&articles)
	return articles
}
