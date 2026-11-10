package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if task.Title == "" {

		writeJson(w, map[string]string{
			"error": "title is required",
		}, http.StatusBadRequest)
		return
	}

	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(apiDate)
	}

	t, err := time.Parse(apiDate, task.Date)

	if err != nil {

		writeJson(w, map[string]string{
			"error": "invalid date",
		}, http.StatusBadRequest)
		return
	}
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil && task.Repeat != "" {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(apiDate)
		} else {
			task.Date = next
		}
	}
	id, err := db.AddTask(&task)

	if err != nil {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	writeJson(w, map[string]int64{"id": id}, http.StatusCreated)
}

func writeJson(w http.ResponseWriter, data any, status int) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(jsonData)

}
