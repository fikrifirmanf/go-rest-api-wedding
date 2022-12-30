package config

import (
	"os"

	model "github.com/fikrifirmanf/go-rest-api-wedding/models"
	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	dbURI := make(chan string, 1)

	if os.Getenv("ENVIRONMENT") != "production" {
		dbURI <- util.GoDotEnv("DB_URI_DEV")
	} else {
		dbURI <- util.GoDotEnv("DB_URI_PROD")
	}

	db, err := gorm.Open(mysql.Open(<-dbURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("ENVIRONMENT") != "production" {
		logrus.Info("Database connection success")
	}

	err = db.AutoMigrate(
		&model.Users{},
		// &model.EntityGuests{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
