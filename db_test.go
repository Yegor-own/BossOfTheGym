package main

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gymboss/graph/model"
	"testing"
)

func TestCreateDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("database.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.TrainingDB{}, &model.GymDB{}, &model.CustomerDB{}, &model.PurchaseDB{})

}
