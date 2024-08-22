package main

import (
	"LinksShortener/internal/handlers"
	"LinksShortener/internal/repositories"
	"LinksShortener/internal/services"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to connect to the database")
	}
	defer db.Close()
	mainRepositories := repositories.InitRepositories(db)
	mainServices := services.InitServices(&mainRepositories.Shortener)
	mainHandlers := handlers.InitHandlers(&mainServices.Shortener)
	r := chi.NewRouter()
	r.Post("/", mainHandlers.Shortener.SetLink)
	r.Get("/{shortLink}", mainHandlers.Shortener.GetLink)
	log.Println("Starting server on: ", os.Getenv("SERVER_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), r)) // Start server

}
