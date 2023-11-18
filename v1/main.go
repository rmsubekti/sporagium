package v1

import (
	"github.com/rmsubekti/sporagium/middleware"
	"github.com/rmsubekti/sporagium/v1/handler"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	var h handler.V1Handler

	v1 := r.PathPrefix("/v1").Subrouter()
	{
		v1.HandleFunc("/register", h.Register).Methods("POST")
		v1.HandleFunc("/login", h.Login).Methods("POST")
	}

	account := v1.PathPrefix("/account").Subrouter()
	account.Use(middleware.JwtAuthMiddleware)
	{
		account.HandleFunc("", h.UserProfile).Methods("GET")
	}

	spora := v1.PathPrefix("/spora").Subrouter()
	spora.Use(middleware.JwtAuthMiddleware)
	{
		spora.HandleFunc("", h.CreateSpora).Methods("POST")
		spora.HandleFunc("/{id}", h.CreateClientSecret).Methods("PATCH")
		spora.HandleFunc("", h.ViewListSpora).Methods("GET")
	}
}
