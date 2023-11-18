package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmsubekti/sporagium/helper"
	"github.com/rmsubekti/sporagium/oauth2server"
	v1 "github.com/rmsubekti/sporagium/v1"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	port := helper.GetEnv("SPORAGIUM_PORT", "80")

	v1.Setup(r)
	oauth2server.Setup(r)

	nuxtDist := http.FileServer(http.Dir("./frontend/.output/public"))
	r.PathPrefix("/").Handler(http.StripPrefix("", nuxtDist))
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "GET", "PATCH", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		// Debug:            true,
	})

	handler := c.Handler(r)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Print(err)
	}
}
