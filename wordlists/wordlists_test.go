package wordlists_test

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/zackslash/sansseed/wordlists"
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
