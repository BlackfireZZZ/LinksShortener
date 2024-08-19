package main

import (
	"LinksShortener/internal/handlers"
	"LinksShortener/internal/repositories"
	"LinksShortener/internal/services"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
	}
	defer db.Close()
	mainRepositories := repositories.InitRepositories(db)
	mainServices := services.InitServices(&mainRepositories.Shortener)
	mainHandlers := handlers.InitHandlers(&mainServices.Shortener)
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/", mainHandlers.Shortener.Shortener)
	})
	log.Println("Starting server on: ", os.Getenv("SERVER_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), r)) // Start server

}
