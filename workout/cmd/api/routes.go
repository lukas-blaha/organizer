package main

import "net/http"

func (app *Config) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /workout/{user}", app.SendNotification)

	return mux
}

func (app *Config) SendNotification(w http.ResponseWriter, r *http.Request) {

}
