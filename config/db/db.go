package db

import (
	"fmt"
	"social_media/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	// Database configuration
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set maximum open connections
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(100)

	// Set maximum idle connections
	sqlDB.SetMaxIdleConns(10)

	return db, nil
}
