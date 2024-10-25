package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct {
	DB *sql.DB
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
