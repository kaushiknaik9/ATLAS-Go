package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kaushiknaik9/ATLAS-Go/internal/response"
)

type Student struct {
	Id    int
	Name  string
	Email string
	Age   int8
}

func Getstudents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students := Student{
			Id:    1,
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Age:   20,
		}
		slog.Info("Getting Student")
		response.SendResponse(w, http.StatusOK, map[string]any{
			"status": "ok",
			"id":     students.Id,
			"name":   students.Name,
		})
	}
}

func GetStudentId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := r.URL.Query().Get("id")

		// id, err := strconv.ParseUint(id_str, 10, 64)
		id, err := strconv.Atoi(id_str)

		if err != nil {
			response.SendErrorResponse(w, http.StatusBadRequest, "Error while Getting Id or Parsing Id")
		}
		response.SendResponse(w, http.StatusOK, map[string]any{
			"status":  "ok",
			"message": "Fetched Id from the URL",
			"data":    id,
		})
	}
}

func CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.SendErrorResponse(w, http.StatusBadRequest, "Error While Reading the Body")
		}
		response.SendResponse(w, http.StatusAccepted, map[string]any{
			"success": "ok",
			"message": "User Creation success",
			"data":    student,
		})
	}
}

func UpdateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.SendErrorResponse(w, http.StatusBadRequest, "Error While Reading the Body")
		}
		response.SendResponse(w, http.StatusAccepted, map[string]any{
			"success": "ok",
			"message": "User Update success",
			"data":    student,
		})
	}
}

func DeleteStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.SendErrorResponse(w, http.StatusBadRequest, "Error While Reading the Body")
		}
		response.SendResponse(w, http.StatusAccepted, map[string]any{
			"success": "ok",
			"message": "User Delete success",
			"data":    student,
		})
	}
}
