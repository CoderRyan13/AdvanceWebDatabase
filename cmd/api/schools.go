// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
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
	// Display the school id
	fmt.Fprintf(w, "show the details for school %d\n", id)
}