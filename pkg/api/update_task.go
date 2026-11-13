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

		writeJson(w, map[string]string{
			"error": "wrong JSON",
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
	_, err := time.Parse(apiDate, task.Date)
	if err != nil {

		writeJson(w, map[string]string{
			"error": "wrong date",
		}, http.StatusBadRequest)
		return
	}
	_, err = NextDate(now, task.Date, task.Repeat)
	if err != nil && task.Repeat != "" {
		writeJson(w, map[string]string{
			"error": "wrong repeat",
		}, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeJson(w, map[string]string{
				"error": "task not found",
			}, http.StatusNotFound)
		} else {
			writeJson(w, map[string]string{
				"error": "internal server error",
			}, http.StatusInternalServerError)
		}
		return
	}

	writeJson(w, map[string]string{}, http.StatusOK)

}
