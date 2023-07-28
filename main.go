package main

import (
	"log"
	api "qdrant/routes"
)

func main() {

	// connection.QdrantDBConn()

	err := api.Init()
	if err != nil {
		log.Fatal("Unable to start the Server.")
	}
}
