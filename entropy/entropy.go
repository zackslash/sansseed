package entropy

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"

	"github.com/zackslash/sansseed/lengths"
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
