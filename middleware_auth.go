package main

import (
	"fmt"
	"net/http"

	"github.com/aiste-i/rssagg/internal/auth"
	"github.com/aiste-i/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, error := auth.GetAPIKey(r.Header)
		if error != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", error))
			return
		}

		user, error := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if error != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %s", error))
			return
		}
		handler(w, r, user)
	}
}
