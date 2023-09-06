package main

import (
	"context"
	"net/http"

	"github.com/oggy107/rss-aggregator/internal/auth"
	"github.com/oggy107/rss-aggregator/respond"
)

const CONFIG_CTX = "config"
const USER_CTX = "user"

// middleware to pass database using context to all handlers
func (cfg *ApiConfig) ApiContext(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), CONFIG_CTX, cfg)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func authorizedOnly(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respond.WithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := r.Context()

		config := ctx.Value("config").(*ApiConfig)
		user, err := config.DB.GetUserByAPIKey(ctx, apiKey)

		if err != nil {
			respond.WithError(w, http.StatusUnauthorized, auth.InvalidApiKey{}.Error())
			return
		}

		userCtx := context.WithValue(ctx, USER_CTX, user)

		next.ServeHTTP(w, r.WithContext(userCtx))
	}

	return http.HandlerFunc(fn)
}
