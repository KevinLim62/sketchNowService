package service

import (
	"context"

	"log"
	"net/http"
	"sketchNow_service/db"
	"sketchNow_service/lib"
	"sketchNow_service/repository"

	"github.com/google/uuid"
)


var ctx = context.Background()

func CreateBoardRoom(w http.ResponseWriter, apiConfig *db.ApiConfig, createBoardRoomInput repository.CreateBoardRoomParams) {


	boardRoom, err := apiConfig.DB.CreateBoardRoom(ctx, createBoardRoomInput)

	if err != nil {
		log.Fatal(err)
	}

	lib.RespondWithJSON(w, 200, boardRoom)
}

func GetAllBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig ) {
	boardRooms, err := apiConfig.DB.GetAllBoardRooms(ctx)

	if err != nil {
		log.Fatal(err)
	}

	lib.RespondWithJSON(w, 200, boardRooms)
}

func GetOneBoardRoom(w http.ResponseWriter, apiConfig *db.ApiConfig, boardRoomId uuid.UUID) {
	boardRoom, err := apiConfig.DB.GetBoardRoomById(ctx, boardRoomId)

	if err != nil {
		log.Fatal(err)
	}

	lib.RespondWithJSON(w, 200, boardRoom)
}

func UpdateBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig, updateBoardRoomInput repository.UpdateBoardRoomByIdParams) {

	boardRoom, err := apiConfig.DB.UpdateBoardRoomById(ctx, updateBoardRoomInput)

	if err != nil {
		log.Fatal(err)
	}

	lib.RespondWithJSON(w, 200, boardRoom)
}

func DeleteBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig, boardRoomId uuid.UUID) {
	err := apiConfig.DB.DeleteBoardRoomById(ctx, boardRoomId)

	if err != nil {
		log.Fatal(err)
	}

	lib.RespondWithJSON(w, 200, "board room deleted successfully")
}