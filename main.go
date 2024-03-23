package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"sketchNow_service/db"
	"sketchNow_service/handler"
	"sketchNow_service/lib"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT environment variable not set")
	}

	hub := lib.NewHub()
	go hub.Run()

	r := chi.NewRouter()
    r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	apiConfig , err := db.ConnectDb();
	if err != nil {
		fmt.Printf("Error connecting to database")
	}

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

	r.Route("/api", func(r chi.Router) {
		r.Mount("/boardRoom", handler.BoardRoomRouter(&apiConfig))

		
		r.Get("/websocket", func(w http.ResponseWriter, r *http.Request) {
			handler.WebsocketHander(hub,w,r)
		})
	})


	http.ListenAndServe(":"+portString, r)
}