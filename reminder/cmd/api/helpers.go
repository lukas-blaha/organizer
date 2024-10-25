package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (app *Config) GetNextId() int {
	var id int

	for _, reminder := range app.Reminders {
		if reminder.Id > id {
			id = reminder.Id
		}
	}

	return id + 1
}

func (e *Entry) GetNextTime() string {
	var next string

	if !e.Done {
		h, m, _ := time.Now().Clock()

		for i := 0; i <= 60; i++ {
			if m < i*e.Repeat {
				if i*e.Repeat == 60 {
					next += fmt.Sprintf("%d 0 0", h+1)
				} else {
					next += fmt.Sprintf("%d %d 0", h, i*e.Repeat)
				}
				break
			}
		}
	}

	if checkLast(next, e.Last) {
		return e.Start
	}

	return next
}

func checkLast(next, last string) bool {
	var ni, li []int

	ns := strings.Split(next, " ")
	ls := strings.Split(last, " ")

	for _, i := range ns {
		n, _ := strconv.Atoi(i)
		ni = append(ni, n)
	}

	for _, i := range ls {
		n, _ := strconv.Atoi(i)
		li = append(li, n)
	}

	if ni[0] > li[0] || ni[0] == li[0] && ni[1] > li[1] {
		return true
	}

	return false
}
