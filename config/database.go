package config

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?%s",
		url.QueryEscape(ENV.DB_USERNAME),
		url.QueryEscape(ENV.DB_PASSWORD),
		ENV.DB_URL,
		url.QueryEscape(ENV.DB_NAME),
		"sslmode=disable&TimeZone=Asia/Jakarta",
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
