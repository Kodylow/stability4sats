package frontend

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"text/template"

	"github.com/kodylow/stability4sats/handlers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new request
	req, err := http.NewRequest("GET", "/v1/engines/list", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function directly
	handlers.EnginesListHandler(rr, req)

	// Decode the list of engines
	var engines []handlers.Engine
	if err := json.NewDecoder(rr.Body).Decode(&engines); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Engines list: ", engines)

	// Execute the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Engines": engines,
	})
	if err != nil {
		log.Println("Error executing template:", err)
	}
}