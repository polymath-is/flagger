package hashutils

import (
	"github.com/mitchellh/hashstructure"
)

// Hashers are resources which have a custom hashing function defined.
// Hash functions are generated by default for solo-kit resources
type Hasher interface {
	Hash() uint64
}

// hash one or more values
// order matters
func HashAll(values ...interface{}) uint64 {
	var hashes []uint64
	for _, v := range values {
		hashes = append(hashes, hashValue(v))
	}
	return hashValue(hashes)
}

func hashValue(val interface{}) uint64 {
	if hasher, ok := val.(Hasher); ok {
		return hasher.Hash()
	}
	h, err := hashstructure.Hash(val, nil)
	if err != nil {
		panic("resource failed to hash: " + err.Error())
	}
	return h
}
