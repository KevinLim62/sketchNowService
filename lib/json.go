package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}){
	data, err := json.Marshal(payload)

	if err != nil{
		log.Printf("Failed to marshal JSON response: %v",payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, statusCode int, msg string){
	if statusCode > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, statusCode, errResponse{
		Error: msg,
	})
}