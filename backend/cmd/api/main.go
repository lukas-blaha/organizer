package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendNotification(w http.ResponseWriter, r *http.Request) {
	var msg string
	url := "https://ntfy.sh/lukebl-workout"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	msg = string(b)

	resp, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(msg)))
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, fmt.Sprintf("%d\n", resp.StatusCode))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /notify", sendNotification)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}
