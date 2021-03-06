package main

import (
	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/internal/db"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func migrate(instance *gorm.DB) {
	instance.AutoMigrate(&entities.User{})
	instance.AutoMigrate(&entities.Expense{})
	instance.AutoMigrate(&entities.Payable{})

	addCheckForEmptyUsername := "ALTER TABLE users ADD CONSTRAINT check_empty_username CHECK (username <> '');"
	instance.Exec(addCheckForEmptyUsername)

	addCheckForEmptyMobileNumber := "ALTER TABLE users ADD CONSTRAINT check_empty_mobile_number CHECK (mobile_number <> '');"
	instance.Exec(addCheckForEmptyMobileNumber)
}

func main() {
	cfg := config.New()

	err := cfg.Validate()

	if err != nil {
		log.Fatal(err)
	}

	dbInstance := db.New(cfg)

	log.Infof("Initializing migration in %s environment...\n", cfg.AppEnv)

	migrate(dbInstance)
}
