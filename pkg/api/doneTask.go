package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "id is required",
		})
		return
	}

	task, err := db.GetTask(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		writeJson(w, map[string]string{
			"error": "Задача не найдена",
		})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
	} else {
		next, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJson(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
		err = db.UpdateDate(next, id)
	}

	writeJson(w, map[string]string{})

}
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	err := db.DeleteTask(id)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		writeJson(w, map[string]string{
			"error": "Задача не найдена",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	writeJson(w, map[string]string{})

}

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		updateTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	}

}
