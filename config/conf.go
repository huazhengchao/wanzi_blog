package config

import (
	"github.com/huazhengchao/wanzi_blog/model"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	// 读取配置文件
	if err := godotenv.Load(".env"); err != nil {
		panic(".env load error")
	}

	// 连接数据库
	model.Database(os.Getenv("dsn"))
}
