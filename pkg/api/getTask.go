package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}
	task, err := db.GetTask(id)
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

	writeJson(w, task)
}
