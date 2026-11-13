package main

import (
	"log"

	"github.com/barzaevhalid/go_final_project/pkg/api"
	"github.com/barzaevhalid/go_final_project/pkg/db"
	"github.com/barzaevhalid/go_final_project/pkg/server"
	_ "modernc.org/sqlite"
)

func main() {
	api.Load()
	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	if err = server.Start(); err != nil {
		log.Println(err)
	}

}
