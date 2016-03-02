package lib

import (
	"math/rand"
	"time"
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
