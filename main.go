package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type Event struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
}

type Scheduler struct {
	events []Event
}

func (s *Scheduler) addEvent(newEvent Event) bool {
	for _, event := range s.events {
		// Check for overlapping events
		if newEvent.StartTime < event.EndTime && newEvent.EndTime > event.StartTime {
			return false // overlap detected
		}
	}
	s.events = append(s.events, newEvent)
	return true // event added successfully
}

func (s *Scheduler) getEvents() []Event {
	return s.events
}

var scheduler Scheduler

// Unified event handler for both GET and POST requests
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var newEvent Event
		if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if added := scheduler.addEvent(newEvent); added {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newEvent)
		} else {
			http.Error(w, "Event overlaps with existing events", http.StatusConflict)
		}
	case http.MethodGet:
		events := scheduler.getEvents()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", eventsHandler) // Only register this route once

	// Set up CORS
	handler := cors.Default().Handler(mux)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
