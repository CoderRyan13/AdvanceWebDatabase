// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"appletree.ryanarmstrong.net/internal/data"
)

// createSchoolHandler for the "Post /v1/schools" endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new school..")
}

// showSchoolHandler for the "Post /v1/schools/:id" endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the School struct containing the ID we extracted
	// from our URL and some sample data
	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Apple Tree",
		Level:     "High School",
		Contact:   "Anna Smith",
		Phone:     "636-3636",
		Address:   "13 Apple street",
		Mode:      []string{"blended", "online"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
