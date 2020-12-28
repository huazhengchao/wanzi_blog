package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Database(dsn string)  {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置表前缀及禁用全局复数
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wanzi_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败")
	}


	DB = db
}
