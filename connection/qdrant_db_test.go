package connection_test

import (
	connection "qdrant/connection"
	"testing"

	pb "github.com/qdrant/go-client/qdrant"
	"github.com/stretchr/testify/assert"
)

var (
	vectorSize uint64 = 4
	distance          = pb.Distance_Dot
)

func TestQdrantDBConn(t *testing.T) {
	expectedAddr := "localhost:6334"
	conn, collectionsClient, ctx, cancel := connection.QdrantDBConn()
	defer cancel()

	assert.NotNil(t, conn)
	assert.NotNil(t, collectionsClient)
	assert.NotNil(t, ctx)

	assert.Equal(t, expectedAddr, conn.Target())

	collectionName := "test_collection_15"
	_, err := collectionsClient.Create(ctx, &pb.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: distance,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &vectorSize,
		},
	})
	assert.NoError(t, err)

	_, err = collectionsClient.Delete(ctx, &pb.DeleteCollection{
		CollectionName: collectionName,
	})
	assert.NoError(t, err)
}
