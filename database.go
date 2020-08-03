package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDatabase(conf *Config) (*gorm.DB, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Name,
		conf.Database.Password,
		conf.Database.SslMode)

	db, err := gorm.Open(conf.Database.Type, connection)
	if err != nil {
		return db, err
	}

	db.AutoMigrate(&User{})
	db.Model(&User{}).AddIndex("idx_user_active", "active")

	return db, err
}
