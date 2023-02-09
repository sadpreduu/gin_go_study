package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mongodb"
)

// Connect to the MongoDB database
func connectDB() *gorm.DB {
	db, err := gorm.Open("mongodb", "mongodb://localhost/your_database_name")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	// Automatically create the reminders collection if it doesn't exist
	db.AutoMigrate(&Reminder{})

	return db
}
