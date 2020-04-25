package lengths

import (
	"errors"
	"strings"
)

// SeedBitLength enumerates bit length required for mnemonic lengths
type SeedBitLength int

const (
	// UnknownWordSeed is used on error
	UnknownWordSeed SeedBitLength = -1

	// TwelveWordSeed 128 bit 12 word seed
	TwelveWordSeed SeedBitLength = 128

	// FiveteenWordSeed 160 bit 15 word seed
	FiveteenWordSeed SeedBitLength = 160

	// EighteenWordSeed 192 bit 18 word seed
	EighteenWordSeed SeedBitLength = 192

	// TwentyoneWordSeed 224 bit 21 word seed
	TwentyoneWordSeed SeedBitLength = 224

	// TwentyfourWordSeed 256 bit 24 word seed
	TwentyfourWordSeed SeedBitLength = 256
)

// ChecksumBitsLen calculates length of the checksum for given seed length
func ChecksumBitsLen(n SeedBitLength) int {
	return int(n) / 32
}

// GetLengthTypeForMnemonic attempts to detect the mnemonic length type given the string
func GetLengthTypeForMnemonic(s string) (SeedBitLength, error) {
	wordCount := len(strings.Split(s, " "))
	switch wordCount {
	case 12:
		return TwelveWordSeed, nil
	case 15:
		return FiveteenWordSeed, nil
	case 18:
		return EighteenWordSeed, nil
	case 21:
		return TwentyoneWordSeed, nil
	case 24:
		return TwentyfourWordSeed, nil
	}

	return UnknownWordSeed, errors.New("Unable to determine seed type")
}
