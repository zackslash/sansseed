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
	"testing"

	"github.com/sanscentral/sansseed/lengths"
)

func TestEntropyGeneration(t *testing.T) {
	l := []lengths.SeedBitLength{lengths.TwelveWordSeed, lengths.FiveteenWordSeed, lengths.EighteenWordSeed, lengths.TwentyoneWordSeed, lengths.TwentyfourWordSeed}
	for _, seedLen := range l {
		cl := lengths.ChecksumBitsLen(seedLen)
		words := (cl + int(seedLen)) / 11
		e, err := GetRandomEntropyBytesWithCheckSum(seedLen)
		if err != nil {
			t.Errorf("Failed to generate %d word seed entropy: %s", words, err.Error())
		}

		if len(e)%11 != 0 {
			t.Errorf(" %d word seed entropy cannot be split into words", words)
		}

		ln := int(seedLen) + lengths.ChecksumBitsLen(seedLen)
		if len(e) != ln {
			t.Errorf("%d seed entropy is invalid length. expected %d got %d", words, ln, len(e))
		}
		t.Logf("%d seed entropy is : %s", words, e)
	}
}

func TestBinaryStringToInt(t *testing.T) {
	bl := []string{
		"100001111100010010001101000010011000011001000011010010100011100110101011000011011011011001010001000110101010101111011100010011010110",
		"000100111110100011111111000010010011010000100110010011010011101110010010011111000100111110001110011001011100111000110110101001000001100001010100101100110011010000111",
		"001000110011010101111010001010100110000101000010100010100100011010111011100011010000010010011111010100101101110000110100111000011111111110111011100101011100111011101001000101000111111101111111111101",
		"111100001010011111110110010010010110010101010001000110100101011100001100100001010001101101001010001010100111000100111101000000110111010111111000110001110110101110101100110100000011101010100110011110010101110000000110001001101111011",
		"011001101110010110001101111000110110110011001111010010001100001110110111010010100001110101110110100001101110011111111011011011110101100001111111111100101011111000000010100010000101000100101101000010010010000011010101100101110011101110010000000111100010010111111110",
	}

	for _, e := range bl {
		s, err := BinaryStringToIntSlice(e)
		if err != nil {
			t.Error(err)
		}
		t.Logf("BinToInt\n%s\n%v\n", e, s)
	}
}
