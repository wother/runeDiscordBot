package workers

import (
	"math/rand"
)

// RandomFromArray returns a random element from a string array.
func RandomFromArray(arr []string, r *rand.Rand) string {
	return arr[r.Intn(len(arr))]
}
