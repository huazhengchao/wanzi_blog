package model

import (
	"time"
)

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Img string `json:"img"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// 文章列表
func (a *Article) List(page int, title string) ([]Article, int64) {
	var articles []Article
	offset := (page - 1) * 10
	var total int64
	DB.Where("title LIKE ?", "%" + title + "%").
	   Offset(offset).
	   Limit(10).
	   Order("created_at desc").
	   Omit("content").
	   Find(&articles).
	   Count(&total)
	return articles, total
}

// 文章详情
func (a *Article) Detail(id int) Article {
	var articles Article
	DB.Where("id = ?", id).First(&articles)
	return articles
}