package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func runServer(router *chi.Mux) {
	const SERVER_NAME = "localhost"

	PORT := os.Getenv("PORT")
	SERVER_ADDR := SERVER_NAME + ":" + PORT

	if PORT == "" {
		log.Fatalln("PORT is not set")
	}

	log.Println("Starting server at:", SERVER_ADDR)

	if serverError := http.ListenAndServe(SERVER_ADDR, router); serverError != nil {
		log.Fatalln("[ERROR]:", serverError)
	}
}

func main() {
	godotenv.Load(".env")

	router := chi.NewRouter()

	handler := &Handler{
		v1: &V1{},
	}

	router.Get("/", handler.root)

	router.Route("/v1", func(v1 chi.Router) {
		v1.Get("/", handler.v1.root)
		v1.Get("/users", handler.v1.users)
	})

	runServer(router)
}
