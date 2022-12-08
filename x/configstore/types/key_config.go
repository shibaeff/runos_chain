package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ConfigKeyPrefix is the prefix to retrieve all Config
	ConfigKeyPrefix = "Config/value/"
)

// ConfigKey returns the store key to retrieve a Config from the index fields
func ConfigKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
