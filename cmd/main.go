package main

import (
	"fmt"
	"net/http"
	"shortLinks/configs"
	"shortLinks/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}
