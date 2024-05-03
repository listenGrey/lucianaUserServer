package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// MySQLClient 创建MySQL客户端
func MySQLClient() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:tinghui0430@tcp(localhost:3306)/schema_luciana"), &gorm.Config{})
	if err != nil {
		log.Printf("无法连接到MySQL")
		return nil
	}
	return db
}
