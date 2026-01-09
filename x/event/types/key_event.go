package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EventKeyPrefix is the prefix to retrieve all Event
	EventKeyPrefix = "Event/value/"
)

// EventKey returns the store key to retrieve a Event from the index fields
func EventKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
