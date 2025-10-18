package app

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() error {
	var err error
	DB, err = gorm.Open(mysql.Open(Conf.Database.Source))
	if err != nil {
		log.Errorf("Failed to connect to MySQL: %v", err)
		return err
	}
	return nil
}
