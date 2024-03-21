package routes

import (
	"sketchNow_service/services"

	"github.com/go-chi/chi/v5"
)

func BoardRoomRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", services.CreateBoardRoom)
	r.Get("/", services.GetAllBoardRoom)
	r.Get("/{boardRoomId}", services.GetOneBoardRoom)
	r.Put("/{boardRoomId}", services.UpdateBoardRoom)
	return r
}

