package api

import (
	"net/http"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	}, http.StatusOK)
}
