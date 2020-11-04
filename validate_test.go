package hashvalidate

import (
	"testing"

	"github.com/tankbusta/hashvalidate/hashes"
)

func TestValidateEmbeddedHashes(t *testing.T) {
	types := hashes.GetTypes()

	for _, hashType := range types {
		t.Run(hashType.Name(), func(subT *testing.T) {
			if err := ValidateHash(hashType.Type(), hashType.Example()); err != nil {
				subT.Fatalf("Failed to validate %s: %s", hashType.Name(), err.Error())
			}
		})
	}
}
