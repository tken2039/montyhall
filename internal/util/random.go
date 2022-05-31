package util

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(to int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(to)
}
