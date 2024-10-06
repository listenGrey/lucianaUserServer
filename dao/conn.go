package dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormDB() (*gorm.DB, error) {
	dsn := "user=me password=tinghui0430 dbname=luciana_user host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // 启用详细日志
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v\n", err)
	}

	return db, nil
}
