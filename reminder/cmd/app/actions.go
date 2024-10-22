package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func (app *Config) Cron() {
	for {
		for i, reminder := range app.Reminders {
			var t string
			h, m, s := time.Now().Clock()

			t = fmt.Sprintf("%d %d %d", h, m, s)

			if reminder.Next == t {
				reminder.ActionByCategory()
				app.Reminders[i].Next = reminder.GetNextTime()
			}
		}
	}
}

func (e *Entry) ActionByCategory() {
	var url string

	if strings.ToLower(e.Category) == "workout" {
		if strings.ToLower(e.User) == "lukas" {
			url = "http://backend:8080/workout"

			_, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(`{"exercise": "pushups", "count": 0}`)))
			if err != nil {
				log.Print(err)
			}

			time.Sleep(time.Second)
		}
	}

}
