package main

import (
	"SJTU-Canteen-Community/internal/app"
	"SJTU-Canteen-Community/internal/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := app.LoadConfig()
	if err != nil {
		log.Errorf("Failed to load config: %v", err)
		return
	}

	DB, err := gorm.Open(mysql.Open(app.Conf.Database.Source))
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
		return
	}

	err = DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Errorf("Failed to migrate database: %v", err)
		return
	}
}
