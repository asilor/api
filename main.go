package main

import (
	"os"
    "net/http"
	"log"

	"github.com/joho/godotenv"

    "std-mongo/database"
	"std-mongo/handlers"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println(".env file not found")
    }

	uri := os.Getenv("MONGO_URI")
	db := os.Getenv("MONGO_DB")

	database.InitDB(uri, db)
	defer database.CloseDB()
    
    mux := http.NewServeMux()

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	handlers.ProductRoutes(mux)

	log.Print("Server started on port http://localhost:8000")
    http.ListenAndServe(":8000", mux)
}
