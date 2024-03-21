package main

import (
	"net/http"

	"sketchNow_service/db"
	"sketchNow_service/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

	db.ConnectDb()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
	

	r.Route("/api", func(r chi.Router) {
		r.Mount("/boardRoom", routes.BoardRoomRouter())
	})

	
	http.ListenAndServe(":5000", r)
}