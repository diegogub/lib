package lib

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	r *rand.Rand
)

func BuildPrefix(ids ...string) string {
	return strings.Join(ids, "")
}

// Prefix should be unique within clusters
func NewLongId(prefix string) string {
	var id string
	var base int = 10 ^ 24
	//random number
	num := base + r.Intn(999999999999999999)
	// ticks
	nsec := uint64(time.Now().UnixNano() / 100)
	id = prefix + strconv.FormatUint(uint64(num), 32) + strconv.FormatUint(nsec, 32)

	return strings.ToUpper(id)
}

func NewShortId(prefix string) string {
	var id string
	// ticks
	nsec := uint64(time.Now().UnixNano() / 100)
	id = prefix + strconv.FormatUint(nsec, 32)
	return strings.ToUpper(id)
}
