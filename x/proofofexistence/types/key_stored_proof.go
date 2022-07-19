package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProofKeyPrefix is the prefix to retrieve all Proof
	ProofKeyPrefix = "Proof/value/"
)

// ProofKey returns the store key to retrieve a Proof from the index fields
func ProofKey(
	hash string,
) []byte {
	var key []byte

	hashBytes := []byte(hash)
	key = append(key, hashBytes...)
	key = append(key, []byte("/")...)

	return key
}
