package service

import (
	"context"

	"net/http"
	"sketchNow_service/db"
	"sketchNow_service/repository"

	"github.com/google/uuid"
)


var ctx = context.Background()

func CreateBoardRoom(w http.ResponseWriter, apiConfig *db.ApiConfig, createBoardRoomInput repository.CreateBoardRoomParams) (repository.Boardroom, error) {
	boardRoom, err := apiConfig.DB.CreateBoardRoom(ctx, createBoardRoomInput)

	if err != nil {
		return boardRoom, err
	}

	return boardRoom, nil
}

func GetAllBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig ) ([]repository.Boardroom, error)  {
	boardRooms, err := apiConfig.DB.GetAllBoardRooms(ctx)

	if err != nil {
		return boardRooms, err
	}

	return boardRooms, nil
}

func GetOneBoardRoom(w http.ResponseWriter, apiConfig *db.ApiConfig, boardRoomId uuid.UUID) (repository.Boardroom, error) {
	boardRoom, err := apiConfig.DB.GetBoardRoomById(ctx, boardRoomId)

	if err != nil {
		return boardRoom, err
	}

	return boardRoom, nil
}

func UpdateBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig, updateBoardRoomInput repository.UpdateBoardRoomByIdParams) (repository.Boardroom, error) {

	boardRoom, err := apiConfig.DB.UpdateBoardRoomById(ctx, updateBoardRoomInput)

	if err != nil {
		return boardRoom, err
	}

	return boardRoom, nil
}

func DeleteBoardRoom(w http.ResponseWriter,  apiConfig *db.ApiConfig, boardRoomId uuid.UUID) (error) {
	err := apiConfig.DB.DeleteBoardRoomById(ctx, boardRoomId)

	if err != nil {
		return err
	}
	
	return nil
}