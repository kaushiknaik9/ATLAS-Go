package routes

import (
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/handlers"
)

func RegisterStudentRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /student", handlers.Getstudents())

	mux.HandleFunc("GET /student/{id}", handlers.GetStudentId())

	mux.HandleFunc("POST /student", handlers.CreateStudent())

	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudent())

	mux.HandleFunc("DELETE /student/{id}", handlers.DeleteStudent())
}
