package database

import (
 "database/sql"
 "fmt"
 _ "github.com/go-sql-driver/mysql"

)

// DBConfig holds configuration details for the database
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}


// InitDB initializes the database connection
func InitDB(config *DBConfig) (*sql.DB, error) {
	// Build connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("连接测试失败: %w", err)
	}

	fmt.Println("连接数据库成功!")
	return db, nil
}