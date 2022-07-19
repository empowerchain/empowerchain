package keeper

var (
	ProofKey = []byte{0x01}
)

// proofStoreKey returns the byte representation of the proof class key
func proofStoreKey(hash string) []byte {
	key := make([]byte, len(ProofKey)+len(hash))
	copy(key, ProofKey)
	copy(key[len(ProofKey):], hash)
	return key
}
