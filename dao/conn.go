package dao

import (
	"log"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

var PostgreConn *sql.DB

func init() {
	var err error
	connStr := "user=" + os.Getenv("POSTGRE_USER") + " dbname=" + os.Getenv("POSTGRE_DB") + " password=" + os.Getenv("POSTGRE_PWD") + " host=" + os.Getenv("POSTGRE_HOST") + " port=" + os.Getenv("POSTGRE_PORT") + " sslmode=disable"
	PostgreConn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 设置连接池参数
	PostgreConn.SetMaxOpenConns(20)   // 最大打开连接数
	PostgreConn.SetMaxIdleConns(10)   // 最大空闲连接数
	PostgreConn.SetConnMaxLifetime(0) // 连接的最大生命周期，0表示无限期
}
