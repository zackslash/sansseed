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

package wordlists_test

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/sanscentral/sansseed/wordlists"
)

// expected value will need to increase when a new language is added
// or TestWordlistAvailability() should fail
const (
	expectedNumberOfWordLists = 8

	// Expected number of words in wordlist (as specified in https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki#Wordlist)
	expectedWordsInList = 2048
)

func TestWordlistAvailability(t *testing.T) {
	l := len(wordlists.GetAll())
	if l != expectedNumberOfWordLists {
		t.Errorf("Wordlist Expected %d languages, got %d", expectedNumberOfWordLists, l)
	}
}

func TestWordlistIntegrity(t *testing.T) {
	l := wordlists.GetAll()

	// Compare hashes to check integrity
	for _, ls := range l {

		n := len(ls.GetList())
		if n != expectedWordsInList {
			t.Errorf("Wordlist %s language integrity check failed due to incorrect count (expecting %d got %d)", ls.GetName(), expectedWordsInList, n)
		}

		for _, w := range ls.GetList() {
			if len(w) == 0 {
				t.Errorf("Wordlist %s language integrity check failed due to a blank entry", ls.GetName())
			}
		}

		// To compare checksum strings are joined with no whitespace
		// integrity hash was generated from BIP document removing all whitespace
		// Example: https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/spanish.txt
		r := strings.Join(ls.GetList(), "")

		h := md5.New()
		h.Write([]byte(r))
		hr := hex.EncodeToString(h.Sum(nil))
		if hr != ls.GetChecksum() {
			t.Errorf("Wordlist %s language MD5 integrity check failed", ls.GetName())
		}
	}
}
