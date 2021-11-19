package my_rand

import (
	"math/rand"
	"time"
)

func GetRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
