package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (app *Config) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /notify", app.Forward)
	mux.HandleFunc("GET /notify/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	return mux
}

func (app *Config) Forward(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(b, &app.Message)
	if err != nil {
		log.Print(err)
	}

	app.SendNotification(w, r)
}

func (app *Config) SendNotification(w http.ResponseWriter, r *http.Request) {
	_, err := http.Post(app.Message.Url, "application/json", bytes.NewBuffer([]byte(app.Message.Msg)))
	if err != nil {
		log.Println(err)
	}
}
