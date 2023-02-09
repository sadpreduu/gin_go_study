package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mongodb"
)

type Reminder struct {
	gorm.Model
	Text string    `gorm:"not null"`
	Time time.Time `gorm:"not null"`
}

func main() {
	// Connect to the MongoDB database
	db, err := gorm.Open("mongodb", "mongodb://localhost/your_database_name")
	if err != nil {
		fmt.Printf("Failed to connect to the database: %s", err)
	}
	defer db.Close()

	// Automatically create the reminders collection if it doesn't exist
	db.AutoMigrate(&Reminder{})

	// Create a new Gin router
	r := gin.Default()

	// Create a GET endpoint to retrieve all reminders
	r.GET("/reminders", func(c *gin.Context) {
		var reminders []Reminder
		db.Find(&reminders)
		c.JSON(200, reminders)
	})

	// Create a POST endpoint to create a new reminder
	r.POST("/reminders", func(c *gin.Context) {
		var reminder Reminder
		if err := c.ShouldBindJSON(&reminder); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&reminder)
		c.JSON(200, reminder)
	})

	// Start the API server
	r.Run(":8000")
}
