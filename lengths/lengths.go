/*
	SansSeed is a BIP39 compatible implementation for generating mnemonic phrases and seed derivation
	Copyright (C) 2018  Sans Central

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
