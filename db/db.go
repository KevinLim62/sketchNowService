package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sketchNow_service/repository"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct{
DB *repository.Queries
}


func ConnectDb() (ApiConfig, error) {
	ctx := context.Background()
	godotenv.Load()
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}


	apiCfg := ApiConfig{
		DB: repository.New(db),
	}


   if err := db.PingContext(ctx); err != nil {
	return apiCfg, err
   }

   fmt.Println("Connected to the database successfully!")

   return apiCfg, nil
}
