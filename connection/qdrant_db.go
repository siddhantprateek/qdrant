package connection

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:6334", "the address to connect to")
)

func QdrantDBConn() (pb.CollectionsClient, context.Context, context.CancelFunc) {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	collectionsClient := pb.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return collectionsClient, ctx, cancel
}

// var QdrantClient pb.CollectionsClient = QdrantDBConn()
