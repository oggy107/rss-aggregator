package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oggy107/rss-aggregator/internal/database"
)

func runServer(router *chi.Mux) {
	const SERVER_NAME = "localhost"

	PORT := os.Getenv("PORT")
	SERVER_ADDR := SERVER_NAME + ":" + PORT

	if PORT == "" {
		logFatal("PORT is not set")
	}

	log.Println("Starting server at:", SERVER_ADDR)

	if serverError := http.ListenAndServe(SERVER_ADDR, router); serverError != nil {
		log.Fatalln("[ERROR]:", serverError)
	}
}

func connectDB() (conn *sql.DB) {
	CONNECTION_STRING := os.Getenv("POSTGRES_CONNECTION")

	if CONNECTION_STRING == "" {
		logFatal("Database connection string is not set")
	}

	conn, err := sql.Open("postgres", CONNECTION_STRING)

	if err != nil {
		logFatal(err.Error())
	}

	return conn
}

type ApiConfig struct {
	DB *database.Queries
}

// middleware to pass database using context to all handlers
func (cfg *ApiConfig) ApiMiddleware(next http.Handler) http.Handler {
	fun := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "config", cfg)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fun)
}

func main() {
	godotenv.Load(".env")

	conn := connectDB()
	defer conn.Close()

	apiConfig := &ApiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(apiConfig.ApiMiddleware)

	handler := &Handler{
		v1: &V1{},
	}

	router.Get("/ping", handler.pong)
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/v1", http.StatusFound)
	// })

	router.Route("/v1", func(v1 chi.Router) {
		// v1.Get("/", handler.v1.root)
		v1.Post("/user", handler.v1.createUser)
	})

	runServer(router)
}
