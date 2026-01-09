package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EventDayKeyPrefix is the prefix to retrieve all EventDay
	EventDayKeyPrefix = "EventDay/value/"
)

// EventDayKey returns the store key to retrieve a EventDay from the index fields
func EventDayKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
