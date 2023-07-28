package main

import (
	"log"
	connection "qdrant/connection"
	api "qdrant/routes"
)

func main() {

	connection.QdrantDBConn()

	err := api.Init()
	if err != nil {
		log.Fatal("Unable to start the Server.")
	}
}
