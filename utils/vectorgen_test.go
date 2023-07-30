package utils_test

import (
	"qdrant/utils"
	"testing"
)

func TestRandomVector(t *testing.T) {
	size := 10
	min := 0.0
	max := 1.0
	vector := utils.RandomVector(size, float32(min), float32(max))
	if len(vector) != size {
		t.Errorf("RandomVector returned a vector of incorrect size. Expected: %d, Got: %d", size, len(vector))
	}
	for _, value := range vector {
		if value < float32(min) || value > float32(max) {
			t.Errorf("RandomVector generated an element out of range. Value: %f, Min: %f, Max: %f", value, min, max)
		}
	}
}
