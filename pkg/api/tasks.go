package api

import (
	"net/http"

	"github.com/barzaevhalid/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

const limit = 50

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJson(w, map[string]string{
			"error": "invalid method",
		}, http.StatusMethodNotAllowed)
	}
	tasks, err := db.Tasks(limit)
	if err != nil {

		writeJson(w, map[string]string{
			"error": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	}, http.StatusOK)
}
