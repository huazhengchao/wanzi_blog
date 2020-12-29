package model

import (
	"time"
)

type ArticleMessage struct {
	Id        int       `json:"id"`
	Pid       int       `json:"pid"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// 文章列表
func (a *ArticleMessage) List(id int) []ArticleMessage {
	var articleMessage []ArticleMessage
	DB.Where("pid = ?", id).
		Order("created_at desc").
		Find(&articleMessage)
	return articleMessage
}

// 文章新增
func (a *ArticleMessage) Add(pid int, name string, content string) bool {
	articleMessage := ArticleMessage{
		Pid: pid,
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
	}
	DB.Create(&articleMessage)
	return true
}
