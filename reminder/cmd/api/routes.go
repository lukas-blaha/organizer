package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (app *Config) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /reminders/create", app.CreateReminder)
	mux.HandleFunc("GET /reminders/list", app.ListReminders)
	mux.HandleFunc("DELETE /reminders/{id}", app.RemoveReminder)

	return mux
}

func (app *Config) CreateReminder(w http.ResponseWriter, r *http.Request) {
	var entry Entry

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read the data: %v\n", err)
	}
	r.Body.Close()

	err = json.Unmarshal(b, &entry)
	if err != nil {
		log.Printf("Cannot parse json payload: %v\n", err)
		return
	}

	if !app.CheckExistance(entry) {
		fmt.Fprintln(w, "The same reminder already set!")
		return
	}

	entry.Id = app.GetNextId()
	entry.Next = entry.GetNextTime()

	app.Reminders = append(app.Reminders, entry)
	err = app.writeToDB(entry)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Fprintf(w, "New reminder #%d added.\n", entry.Id)
}

func (app *Config) ListReminders(w http.ResponseWriter, r *http.Request) {
	var output string
	for _, reminder := range app.Reminders {
		b, err := json.Marshal(reminder)
		if err != nil {
			log.Print(err)
		}
		if len(output) != 0 {
			output = output + "," + string(b)
		} else {
			output = string(b)
		}
	}

	fmt.Fprintln(w, "["+output+"]")
}

func (app *Config) CheckExistance(e Entry) bool {
	for _, entry := range app.Reminders {
		if entry.User == e.User && entry.Category == e.Category && entry.Start == e.Start {
			return false
		}
	}
	return true
}

func (app *Config) RemoveReminder(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("Cannot parse ID: %s\n", strId)
	}

	for i, reminder := range app.Reminders {
		if reminder.Id == id {
			app.Reminders = append(app.Reminders[:i], app.Reminders[i+1:]...)
			err = app.removeFromDB(reminder)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Fprintf(w, "Reminder #%d removed!\n", id)
		}
	}
}
