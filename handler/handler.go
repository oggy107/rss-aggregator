package handler

import (
	"net/http"

	"github.com/oggy107/rss-aggregator/respond"
)

type v1 struct{}

func GetV1() v1 {
	return v1{}
}

func Pong(w http.ResponseWriter, r *http.Request) {
	respond.WithJson(w, 200, map[string]string{"data": "pong"})
}
