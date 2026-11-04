package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "wrong JSON",
		})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "title is required",
		})
		return
	}

	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(apiDate)
	}
	_, err := time.Parse(apiDate, task.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "wrong date",
		})
		return
	}
	_, err = NextDate(now, task.Date, task.Repeat)
	if err != nil && task.Repeat != "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "wrong repeat",
		})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, map[string]string{})

}
