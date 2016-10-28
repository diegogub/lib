package lib

import (
	"github.com/rs/xid"
	"strings"
)

func BuildPrefix(ids ...string) string {
	return strings.Join(ids, "")
}

// Prefix should be unique within clusters
func NewLongId(prefix string) string {
	var id string
	//var base int64 = 10 ^ 6
	guid := xid.New()

	id = guid.String()
	/*
		//random number
		num := base + r.Int63n(999999999)
		// ticks
		nsec := uint64(time.Now().UnixNano() / 100)
		id = prefix + strconv.FormatUint(uint64(num), 32) + strconv.FormatUint(nsec, 32)
	*/

	return id
}

func NewShortId(prefix string) string {
	var id string
	id = xid.New().String()
	return id
}
