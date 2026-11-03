package server

import (
	"net/http"

	"github.com/barzaevhalid/go_final_project/pkg/api"
)

func Start() error {
	mux := http.NewServeMux()

	api.Init(mux)

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	return http.ListenAndServe(":7540", mux)

}
