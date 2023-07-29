package utils

import (
	"math/rand"
	"time"
)

func RandomVector(size int, min, max float32) []float32 {
	rand.NewSource(time.Now().UnixNano())
	vector := make([]float32, size)
	for i := 0; i < size; i++ {
		randomValue := min + rand.Float32()*(max-min)
		flooredValue := float32(int(randomValue*100)) / 100
		vector[i] = flooredValue
	}
	return vector
}
