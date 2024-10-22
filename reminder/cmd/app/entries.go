package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

type Entry struct {
	Id       int    `json:"id"`
	Start    string `json:"start"`
	Repeat   int    `json:"repeat"`
	Next     string `json:"next"`
	Last     string `json:"last"`
	Category string `json:"category"`
	User     string `json:"user"`
	Done     bool   `json:"-"`
}

type Entries []Entry

func (app *Config) loadSavedData() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, start, repeat, last, category, name FROM reminders`

	log.Println("Loading data...")

	rows, err := app.DB.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var entry Entry
		err := rows.Scan(
			&entry.Id,
			&entry.Start,
			&entry.Repeat,
			&entry.Last,
			&entry.Category,
			&entry.User,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return err
		}

		entry.Done = false
		entry.Next = entry.GetNextTime()

		fmt.Println("Entry: ", entry)
		app.Reminders = append(app.Reminders, entry)
	}

	log.Println("All data loaded!")
	return nil
}

func (app *Config) writeToDB(e Entry) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO reminders (id, start, repeat, last, category, name) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := app.DB.ExecContext(ctx, query, e.Id, e.Start, e.Repeat, e.Last, e.Category, e.User)
	if err != nil {
		return err
	}

	log.Printf("New reminder added for user %s: %s\n", e.User, e.Category)
	return nil
}

func (app *Config) removeFromDB(e Entry) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM reminders where id = $1`

	_, err := app.DB.ExecContext(ctx, query, e.Id)
	if err != nil {
		return err
	}

	log.Printf("Reminder #%d removed!\n", e.Id)
	return nil
}
