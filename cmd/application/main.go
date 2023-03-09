package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Reminder struct {
	ID        int
	Message   string
	Timestamp time.Time
}

var reminders []Reminder

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define the routes for the API
	router.HandleFunc("/reminders", GetReminders).Methods("GET")
	router.HandleFunc("/reminders", CreateReminder).Methods("POST")
	router.HandleFunc("/reminders/{id}", GetReminder).Methods("GET")
	router.HandleFunc("/reminders/{id}", UpdateReminder).Methods("PUT")
	router.HandleFunc("/reminders/{id}", DeleteReminder).Methods("DELETE")

	// Start the HTTP server
	fmt.Println("Starting the server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetReminders retrieves a list of all reminders
func GetReminders(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	w.Write([]byte("Getting all reminders"))
}

// CreateReminder creates a new reminder
func CreateReminder(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	w.Write([]byte("Creating a new reminder"))
}

// GetReminder retrieves the details of a specific reminder
func GetReminder(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	w.Write([]byte("Getting a reminder"))
}

// UpdateReminder updates the details of a specific reminder
func UpdateReminder(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	w.Write([]byte("Updating a reminder"))
}

// DeleteReminder deletes a specific reminder
func DeleteReminder(w http.ResponseWriter, r *http.Request) {
	// Write the response to the client
	w.Write([]byte("Deleting a reminder"))
}
