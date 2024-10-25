package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct {
	Message Message
}

func main() {
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
