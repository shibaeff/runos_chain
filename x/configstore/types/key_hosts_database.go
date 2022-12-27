package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// HostsDatabaseKeyPrefix is the prefix to retrieve all HostsDatabase
	HostsDatabaseKeyPrefix = "HostsDatabase/value/"
)

// HostsDatabaseKey returns the store key to retrieve a HostsDatabase from the index fields
func HostsDatabaseKey(
	dpid string,
	mac string,
) []byte {
	var key []byte

	dpidBytes := []byte(dpid)
	key = append(key, dpidBytes...)
	key = append(key, []byte("/")...)

	macBytes := []byte(mac)
	key = append(key, macBytes...)
	key = append(key, []byte("/")...)

	return key
}
