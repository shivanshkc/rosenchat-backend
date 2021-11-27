package hashutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hex provides the hex encoded SHA256 of the input.
func SHA256Hex(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
