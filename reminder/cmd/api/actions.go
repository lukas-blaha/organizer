package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
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
	var addr string
	var url string

	if strings.ToLower(e.Category) == "workout" {
		addr = os.Getenv("WORKOUT_URL")

		if addr[len(addr)-1] == '/' {
			url = addr + "workout"
		} else {
			url = addr + "/workout"
		}

		url += fmt.Sprintf("/%s", strings.ToLower(e.User))

		_, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte("all")))
		if err != nil {
			log.Print(err)
		}

		time.Sleep(time.Second)
	}

}
