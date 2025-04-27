package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 获取数据库配置
	dbConfig := AppConfig.Database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// 验证连接
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	fmt.Println("Database connection established.")
}