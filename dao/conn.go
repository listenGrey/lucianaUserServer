package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// MySQLClient 创建MySQL客户端
func MySQLClient() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:200004@tcp(localhost:3306)/luciana_user_db"), &gorm.Config{})
	if err != nil {
		log.Printf("无法连接到MySQL")
		return nil
	}
	return db
}
