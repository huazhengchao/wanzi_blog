package model

import (
	"time"
)

type Article struct {
	Id int `json:"id"`
	Pid int `json:"pid"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Img string `json:"img"`
	Content string `json:"content"`
	View int `json:"view"`
	CreatedAt time.Time `json:"created_at"`
}

// 文章列表
func (a *Article) List(page int, title string) ([]Article) {
	var articles []Article
	offset := (page - 1) * 10
	DB.Where("title LIKE ?", "%" + title + "%").
	   Offset(offset).
	   Limit(5).
	   Order("created_at desc").
	   Omit("content").
	   Find(&articles)
	return articles
}

// 文章详情
func (a *Article) Detail(id int) Article {
	var articles Article
	DB.Where("id = ?", id).First(&articles)
	return articles
}