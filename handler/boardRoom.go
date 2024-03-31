package handler

import (
	"encoding/json"
	"net/http"
	"sketchNow_service/db"
	"sketchNow_service/lib"
	"sketchNow_service/repository"
	"sketchNow_service/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func BoardRoomRouter(apiConfig *db.ApiConfig) chi.Router {
	c := chi.NewRouter()

	c.Post("/", func(w http.ResponseWriter, r *http.Request) {
		type reqBody struct {
			Name string `json:"name"`
		}

		var body reqBody
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&body); err != nil {
			lib.RespondWithError(w, 400, "Error parsing request body")
			return // Return to avoid executing the success response code
		}

		payload := repository.CreateBoardRoomParams{
			ID: uuid.New(),
			Name: body.Name,
		}

		result, err := service.CreateBoardRoom(w, apiConfig, payload)
		if err != nil {
			lib.RespondWithError(w, 400, err.Error())
		}
		lib.RespondWithJSON(w, 200, result)
	})

	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := service.GetAllBoardRoom(w, apiConfig)
		if err != nil {
			lib.RespondWithError(w, 400, err.Error())
		}
		lib.RespondWithJSON(w, 200, result)
	})

	c.Get("/{boardRoomId}", func(w http.ResponseWriter, r *http.Request) {

		boardRoomId, err := uuid.Parse(chi.URLParam(r ,"boardRoomId"))
		if err != nil {
			// Handle error: Invalid UUID format
			lib.RespondWithError(w, 400, "Invalid boardRoomId")
			return
		}

		result, err := service.GetOneBoardRoom(w, apiConfig, boardRoomId)
		if err != nil {
			lib.RespondWithError(w, 400, err.Error())
		}
		lib.RespondWithJSON(w, 200, result)
	})

	c.Put("/{boardRoomId}", func(w http.ResponseWriter, r *http.Request) {

		type reqBody struct {
			Name string `json:"name"`
		}

		boardRoomId, err := uuid.Parse(chi.URLParam(r ,"boardRoomId"))
		if err != nil {
			// Handle error: Invalid UUID format
			lib.RespondWithError(w, 400, "Invalid boardRoomId")
			return
		}

		var body reqBody
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&body); err != nil {
			lib.RespondWithError(w, 400, "Error parsing request body")
			return // Return to avoid executing the success response code
		}

		payload := repository.UpdateBoardRoomByIdParams{
			ID: boardRoomId,
			Name: body.Name,
		}

		result, err := service.UpdateBoardRoom(w, apiConfig, payload)
		if err != nil {
			lib.RespondWithError(w, 400, err.Error())
		}
		lib.RespondWithJSON(w, 200, result)
	})
	c.Delete("/{boardRoomId}", func(w http.ResponseWriter, r *http.Request) {

		boardRoomId, err := uuid.Parse(chi.URLParam(r ,"boardRoomId"))
		if err != nil {
			// Handle error: Invalid UUID format
			lib.RespondWithError(w, 400, "Invalid boardRoomId")
			return
		}

		err = service.DeleteBoardRoom(w, apiConfig, boardRoomId)
		if err != nil {
			lib.RespondWithError(w, 400, err.Error())
		}
		lib.RespondWithJSON(w, 200, "Successfully deleted boardRoom")
	})
	return c
}

