package main

import (
	"log"

	"github.com/barzaevhalid/go_final_project/pkg/db"
	"github.com/barzaevhalid/go_final_project/pkg/server"
	_ "modernc.org/sqlite"
)

func main() {
	err := db.Init("scheduler.db")

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Start())
	db.DB.Close()
}
