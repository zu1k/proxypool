package database

import (
	"fmt"
	"os"

	"github.com/zu1k/proxypool/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func connect() (err error) {
	dsn := "user=proxypool password=proxypool dbname=proxypool port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if url := config.Config.DatabaseUrl; url != "" {
		dsn = url
	}
	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		fmt.Println("DB connect success: ", DB.Name())
	}
	return
}
