package security

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashEmail(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
