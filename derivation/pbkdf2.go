package derivation

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

var (
	salt       = "mnemonic"
	iterations = 2048
	keyLength  = 64
)

// DeriveSeedFromMnemonic turns a mnemonic into a 512bit seed
func DeriveSeedFromMnemonic(mnemonic, optionalPassword string) []byte {
	s := salt + optionalPassword
	return pbkdf2.Key([]byte(mnemonic), []byte(s), iterations, keyLength, sha512.New)
}
