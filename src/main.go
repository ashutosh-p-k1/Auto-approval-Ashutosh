package main

import (
	"database/sql"
	"log"
)

func main() {
	log.Println("Auto approval server: starting up ...")
	log.Println("......................................")
	log.Println("Auto policy engine server started")
	log.Print("Listening on Port 8080... call /git and /db and approve to begin")

	if db == nil {
		db, fn, dir, err = createDB()
	}
	if err != nil {
		log.Fatal("Recieved an error", err)
	}
	err = nil
	err = checkOut(db, fn, dir)
	CheckIfError(err)

}

var db *sql.DB
var fn, dir string
