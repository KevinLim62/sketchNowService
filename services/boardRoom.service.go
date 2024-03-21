package services

import "net/http"

func CreateBoardRoom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Board Room"))
}

func GetAllBoardRoom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get All Board Room"))
}

func GetOneBoardRoom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Grt Board Room"))
}

func UpdateBoardRoom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Board Room"))
}