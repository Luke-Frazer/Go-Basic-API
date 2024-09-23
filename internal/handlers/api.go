package handlers

import (
	"github.com/Luke-Frazer/Personal-Projects/tree/main/Go_Basic_API/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// capital H means this can be imported into other packages
func Handler(r *chi.Mux) {
	// global middleware - applied all the time to all endpoints
	// ensures all ending slashes are ignored (ex: https://localhost:8000/my/path/ -> https://localhost:8000/my/path)
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// middleware for /account route
		router.Use(middleware.Authorization) // all requests to /account must pass through authorization function first

		router.Get("/coins", GetCoinBalance) // requests to /coins endpoint call the GetCoinBalance function
	})
}
