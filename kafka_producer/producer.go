package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Payload struct {
	ID             int    `json:"id"`
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

	topic := "payloads"
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := [...]string{"Granada", "Bilbao", "Valencia", "Seville", "Barcelona", "Madrid"}
	// items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10000; n++ {

		key := users[rand.Intn(len(users))]
		payload := Payload{
			ID:             n + 1,
			City:           key,
			Location:       "Spain",
			CollectionName: "test_collection",
		}

		value, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v\n", err)
			continue
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          value,
		}, nil)
	}
	p.Flush(15 * 1000)
	p.Close()
}
