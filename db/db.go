package db

import (
	"log"

	"github.com/nathangds/order-api/helpers"
	"github.com/nathangds/order-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := helpers.GetEnvVariable("DB_URL")
	driver := helpers.GetEnvVariable("DB_DRIVER")
	var dialector gorm.Dialector

	switch driver {

	case "postgres":
		dialector = postgres.Open(dbURL)
		break
	default:
		log.Fatalln("Invalid database driver")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	db.AutoMigrate(models.Category{})

	return db
}
