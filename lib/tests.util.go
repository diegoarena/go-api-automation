package lib

import (
	"math/rand"
	"time"
)

//returns a random number between specified range
func GetRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
