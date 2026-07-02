package response

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"message"`
}

func SendResponse(w http.ResponseWriter, statusCode int, message interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(message)
}

func SendErrorResponse(w http.ResponseWriter, statuscode int, errorMessage string) ErrResponse {
	return ErrResponse{
		StatusCode:   statuscode,
		ErrorMessage: errorMessage,
	}
}

func ValidationError() {

}

func InternalServerError() {

}
