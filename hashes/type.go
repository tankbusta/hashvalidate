package hashes

import (
	"fmt"
	"sync"

	"github.com/tankbusta/hashvalidate/tokenizer"
)

var (
	backends   = make(map[int]IHashType)
	backendsmu = sync.RWMutex{}
)

// IHashType describes the APIs available for each hash type
type IHashType interface {
	// Name returns a string indicating the type of hash this is
	Name() string

	// Example provides a string that can be used to indicate the expected value of this hash type
	Example() string

	// Type returns the unique hashcat identifier for this hash type
	Type() int

	// Tokens returns the format of expected tokens from this hash type. It includes definitions on how to
	// validate each token as well
	Tokens() []tokenizer.Token
}

// Register loads a new hash type into the backend
//
// Care should be taken while using this API as it's designed to run
// during startup and can panic if a hash type already exists (or a collision on hash type IDs)
func Register(hashTypeID int, hashType IHashType) {
	backendsmu.Lock()
	defer backendsmu.Unlock()

	if hashType == nil {
		panic("cannot register a nil hashType")
	}

	if _, exists := backends[hashTypeID]; exists {
		panic(fmt.Sprintf("hash type %d already exists", hashTypeID))
	}

	backends[hashTypeID] = hashType
}

// Open returns the hash type if it exists
func Open(hashTypeID int) (IHashType, error) {
	backendsmu.RLock()
	defer backendsmu.RUnlock()

	if validator, ok := backends[hashTypeID]; ok {
		return validator, nil
	}

	return nil, fmt.Errorf("Hash Type %d does not exist", hashTypeID)
}

// GetTypes returns a list of hash types registered in the backend
func GetTypes() []IHashType {
	backendsmu.RLock()
	defer backendsmu.RUnlock()

	out := make([]IHashType, len(backends))
	i := 0

	for _, hashType := range backends {
		out[i] = hashType
		i++
	}

	return out
}
