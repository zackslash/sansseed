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

package entropy

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"

	"github.com/sanscentral/sansseed/lengths"
)

// GetRandomEntropyBytesWithCheckSum returns securely generated random bits with sha256 suffix
// returns error if the secure random number generator fails to function correctly
func GetRandomEntropyBytesWithCheckSum(n lengths.SeedBitLength) (string, error) {
	entropyByteLen := int(n / 8)
	b := make([]byte, entropyByteLen)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	fentbits := byteSliceToBits(b)
	cb, err := CheckSumBits(b, n)
	if err != nil {
		return "", err
	}
	rbits := fentbits + cb
	return rbits, nil
}

// BinaryStringToIntSlice splits a binary string every 11 bits and returns the resulting integers for each split
func BinaryStringToIntSlice(e string) ([]int, error) {
	if len(e)%11 != 0 {
		return []int{}, errors.New("Failed to split string every 11 bits")
	}

	r := []int{}
	l := len(e) / 11
	for index := 0; index < l; index++ {
		p := 11 * index
		s := e[p : p+11]

		i, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			return []int{}, err
		}
		r = append(r, int(i))
	}

	return r, nil
}

// CheckSumBits returns the first n bytes by seedlength of the SHA256 hash for given bytes
func CheckSumBits(b []byte, n lengths.SeedBitLength) (string, error) {
	c := sha256.Sum256(b)
	rb := lengths.ChecksumBitsLen(n)
	cb := ""
	for _, cby := range c {
		bi := fmt.Sprintf("%08b", cby)
		if rb >= 8 {
			cb += bi
			rb -= 8
		} else {
			// Use beginning
			ln := bi[:rb]
			cb += ln
			break
		}
	}
	return cb, nil
}

// byteSliceToBits returns binary representation of data in a string
func byteSliceToBits(n []byte) string {
	r := ""
	for _, b := range n {
		i := fmt.Sprintf("%08b", b)
		r += i
	}
	return r
}
