package workers

import (
	"math/rand"
	"time"
)

// RandomFromArray returns a random element from a string array.
func RandomFromArray(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	return arr[rand.Intn(len(arr))]
}
