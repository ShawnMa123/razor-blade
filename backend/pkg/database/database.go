package database

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dbPath string) (*gorm.DB, error) {
	// 在Windows开发环境中，如果CGO不可用，使用内存数据库
	var dsn string
	if dbPath == "" || dbPath == ":memory:" {
		dsn = ":memory:"
	} else {
		// 确保数据库目录存在
		dir := filepath.Dir(dbPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create database directory: %w", err)
		}
		dsn = dbPath
	}

	// 连接数据库，添加pragma参数以支持CGO禁用的情况
	db, err := gorm.Open(sqlite.Open(dsn+"?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		// 如果文件数据库失败，尝试使用内存数据库
		if dsn != ":memory:" {
			fmt.Printf("Warning: Failed to open file database, falling back to memory database: %v\n", err)
			db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		}
		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}
	}

	return db, nil
}