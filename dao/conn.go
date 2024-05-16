package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"lucianaUserServer/conf"
)

// MySQLClient 创建MySQL客户端
func MySQLClient() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.DBUser+":"+conf.DBPwd+"@tcp("+conf.DBAddress+")/"+conf.Database), &gorm.Config{})
	if err != nil {
		log.Printf("无法连接到MySQL")
		return nil
	}
	return db
}
