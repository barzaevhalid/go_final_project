package server

import (
	"net/http"

	"github.com/barzaevhalid/go_final_project/pkg/api"
)

const port = "7540"

func Start() error {
	mux := http.NewServeMux()

	api.Init(mux)

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	return http.ListenAndServe(":"+port, mux)

}
