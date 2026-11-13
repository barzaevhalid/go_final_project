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
		writeJson(w, map[string]string{
			"error": "Не указан идентификатор",
		}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)

	if errors.Is(err, sql.ErrNoRows) {
		writeJson(w, map[string]string{
			"error": "Задача не найдена",
		}, http.StatusNotFound)
		return
	}

	if err != nil {
		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	writeJson(w, task, http.StatusOK)
}
