package main

import (
	"context"
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/auth"
	"github.com/oggy107/rss-aggregator/respond"
)

func getApiContext(cfg *config.ApiConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), config.CONFIG_CTX, cfg)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func authorizedOnly(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respond.WithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := r.Context()

		apiConfig := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)
		user, err := apiConfig.DB.GetUserByAPIKey(ctx, apiKey)

		if err != nil {
			respond.WithError(w, http.StatusUnauthorized, auth.InvalidApiKey{}.Error())
			return
		}

		userCtx := context.WithValue(ctx, config.USER_CTX, user)

		next.ServeHTTP(w, r.WithContext(userCtx))
	}

	return http.HandlerFunc(fn)
}
