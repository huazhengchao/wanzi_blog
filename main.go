package main

import (
	"github.com/huazhengchao/wanzi_blog/config"
	"github.com/huazhengchao/wanzi_blog/router"
)

func main() {
	// 初始化配置
	config.Init()

	r := router.NewRouter()

	r.Run(":8585")
}

