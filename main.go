package main

import (
	"log"
	api "qdrant/routes"
)

func main() {

	err := api.Init()
	if err != nil {
		log.Fatal("Unable to start the Server.")
	}
}
