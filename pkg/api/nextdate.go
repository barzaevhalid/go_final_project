package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", fmt.Errorf("Repeat is empty")
	}
	date, err := time.Parse("20060102", dstart)

	if err != nil {
		return "", err
	}
	if repeat == "y" {

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(apiDate), nil

	}
	parts := strings.Fields(repeat)

	if len(parts) != 2 || parts[0] != "d" {
		return "", fmt.Errorf("указан неверный формат repeat")
	}
	partRepeat, err := strconv.Atoi(parts[1])

	if err != nil {
		return "", err
	}
	if partRepeat > 400 {
		return "", fmt.Errorf("Repeat is too big")
	}

	interval, err := strconv.Atoi(parts[1])

	if err != nil {
		return "", err
	}
	for {
		date = date.AddDate(0, 0, interval)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format("20060102"), nil

}

func afterNow(date, now time.Time) bool {
	dateStr := date.Format(apiDate)
	nowStr := now.Format(apiDate)

	return dateStr > nowStr
}

func nextDayHandler(res http.ResponseWriter, r *http.Request) {
	userDate := r.FormValue("date")
	repeat := r.FormValue("repeat")
	now := r.FormValue("now")

	date, err := time.Parse(apiDate, now)

	data, err := NextDate(date, userDate, repeat)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.Write([]byte(data))

}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}
