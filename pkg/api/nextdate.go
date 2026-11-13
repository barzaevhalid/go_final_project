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
	if partRepeat < 1 || partRepeat > 400 {
		return "", fmt.Errorf("Repeat must be between 1 and 400")
	}

	for {
		date = date.AddDate(0, 0, partRepeat)
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

func nextDayHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJson(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	userDate := r.FormValue("date")
	repeat := r.FormValue("repeat")
	now := r.FormValue("now")

	var date time.Time

	if now == "" {
		date = time.Now()
	} else {
		parsedDate, err := time.Parse(apiDate, now)

		if err != nil {
			writeJson(w, "wrong now format", http.StatusBadRequest)
			return
		}
		date = parsedDate
	}

	nextdate, err := NextDate(date, userDate, repeat)
	if err != nil {
		writeJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeText(w, nextdate, http.StatusOK)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		writeJson(w, "method not allowd", http.StatusMethodNotAllowed)
	}

}

func writeText(w http.ResponseWriter, data string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(data))
}
