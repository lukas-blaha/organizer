package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var count = 10

type Workout struct {
	Exercise string `json:"exercise"`
	Count    int    `json:"count"`
}

func sendNotification(w http.ResponseWriter, r *http.Request) {
	var msg string
	url := "https://ntfy.sh/lukebl-notifications"

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

func workout(w http.ResponseWriter, r *http.Request) {
	var msg string
	wrkt := Workout{}

	url := "https://ntfy.sh/lukebl-workout"

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &wrkt)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(wrkt)

	count -= wrkt.Count
	if count > 0 {
		msg = fmt.Sprintf("%d pushups to be done", count)
	} else {
		msg = "You're done for today!"
	}

	resp, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(msg)))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, fmt.Sprintf("%d\n", count))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /notify", sendNotification)
	router.HandleFunc("POST /workout", workout)
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println(err)
	}
}
