package main

import (
	"log"
	"os"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}
	log.Print("Received an error: ", err)

	os.Exit(1)
}
