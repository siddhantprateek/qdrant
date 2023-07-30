package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "qdrant/connection"
	"qdrant/utils"

	pb "github.com/qdrant/go-client/qdrant"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type VectorPayload struct {
	Id             int    `json:"id"`
	City           string `json:"city"`
	Location       string `json:"location"`
	CollectionName string `json:"collectionName"`
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}

	configFile := os.Args[1]
	conf := ReadConfig(configFile)
	conf["group.id"] = "kafka-qdapi-consumer"
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "payloads"
	_ = c.SubscribeTopics([]string{topic}, nil)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}

			// Decode the ev.Value into VectorPayload
			newVectorPayload := new(VectorPayload)
			err = json.Unmarshal(ev.Value, newVectorPayload)
			if err != nil {
				fmt.Println("Error decoding message:", err)
				continue
			}

			waitUpsert := true
			upsertPoints := []*pb.PointStruct{
				{
					Id: &pb.PointId{
						PointIdOptions: &pb.PointId_Num{Num: uint64(newVectorPayload.Id)},
					},
					Vectors: &pb.Vectors{
						VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: utils.RandomVector(4, 0.0, 1.0)}},
					},
					Payload: map[string]*pb.Value{
						"city": {
							Kind: &pb.Value_StringValue{StringValue: newVectorPayload.City},
						},
						"location": {
							Kind: &pb.Value_StringValue{StringValue: newVectorPayload.Location},
						},
					},
				},
			}
			conn, collections_client, ctx, cancel := db.QdrantDBConn()
			defer cancel()
			_ = collections_client

			pointsClient := pb.NewPointsClient(conn)
			resp, err := pointsClient.Upsert(ctx, &pb.UpsertPoints{
				CollectionName: newVectorPayload.CollectionName,
				Wait:           &waitUpsert,
				Points:         upsertPoints,
			})
			if err != nil {
				fmt.Println("Unable to upsert data.")
			}
			fmt.Println("Consumed event: ", resp.Result)
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

		}
	}

	c.Close()

}
