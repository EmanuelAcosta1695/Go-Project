package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting application...")

	store, err := NewPostgresStorage()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIserver(":3000", store)
	server.Run()
}
