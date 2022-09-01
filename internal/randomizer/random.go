package randomizer

import (
	"math/rand"
	"time"
)

func GetRandomNumber(min uint8, max uint8) uint8 {
	// generate a seed using time in nano to always generate a different number
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(int((max-min)+min))) + 1
}
