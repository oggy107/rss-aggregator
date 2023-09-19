package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/handler"
	"github.com/oggy107/rss-aggregator/utils"
)

func runServer(router *chi.Mux) {
	const SERVER_NAME = "localhost"

	PORT := os.Getenv("PORT")
	SERVER_ADDR := SERVER_NAME + ":" + PORT

	if PORT == "" {
		utils.LogFatal("PORT is not set")
	}

	log.Println("Starting server at:", SERVER_ADDR)

	if serverError := http.ListenAndServe(SERVER_ADDR, router); serverError != nil {
		log.Fatalln("[ERROR]:", serverError)
	}
}

func main() {
	godotenv.Load(".env")

	apiConfig := config.Init()

	router := chi.NewRouter()

	router.Use(getApiContext(apiConfig))

	v1Handler := handler.GetV1()

	router.Get("/ping", handler.Pong)

	router.Route("/v1", func(v1 chi.Router) {
		v1.Post("/user", v1Handler.CreateUser)
		v1.Get("/all_feeds", v1Handler.GetAllFeeds)

		// authonly routes
		v1.Group(func(v1Auth chi.Router) {
			v1Auth.Use(authorizedOnly)
			v1Auth.Get("/user", v1Handler.GetUser)
			v1Auth.Post("/feed", v1Handler.CreateFeed)
			v1Auth.Get("/feed/{feed_id}", v1Handler.GetFeed)
			v1Auth.Get("/feeds", v1Handler.GetFeeds)
			v1Auth.Post("/feed_follows", v1Handler.CreateFeedFollows)
			v1Auth.Get("/feed_follows", v1Handler.GetFeedFollows)
		})
	})

	runServer(router)
}
