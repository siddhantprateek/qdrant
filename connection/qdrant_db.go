package connection

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr              = flag.String("addr", "localhost:6334", "the address to connect to")
	vectorSize uint64 = 4
	distance          = pb.Distance_Dot
)

func QdrantDBConn() (
	*grpc.ClientConn,
	pb.CollectionsClient,
	context.Context,
	context.CancelFunc) {
	flag.Parse()

	qdrantAddr := os.Getenv("QDRANT_ADDR")
	if qdrantAddr != "" {
		addr = &qdrantAddr
	}

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	collectionsClient := pb.NewCollectionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return conn, collectionsClient, ctx, cancel
}

// var QdrantClient pb.CollectionsClient = QdrantDBConn()
func CreateQdCollection(collectionName string) (string, error) {
	_, collections_client, ctx, cancel := QdrantDBConn()
	defer cancel()
	var defaultSegmentNumber uint64 = 2
	_, err := collections_client.Create(ctx, &pb.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: distance,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Collection %s created", collectionName), nil
}

// Create keyword field index
func CreateFieldIndex(collectionName, fieldIndexName string) (string, error) {
	conn, collections_client, ctx, cancel := QdrantDBConn()
	defer cancel()
	_ = collections_client

	pointsClient := pb.NewPointsClient(conn)

	fieldIndexType := pb.FieldType_FieldTypeKeyword
	_, err := pointsClient.CreateFieldIndex(ctx, &pb.CreateFieldIndexCollection{
		CollectionName: collectionName,
		FieldName:      fieldIndexName,
		FieldType:      &fieldIndexType,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Field index for field  %s created", fieldIndexName), nil
}
