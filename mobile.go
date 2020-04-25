package sansseed

import (
	"encoding/hex"
	"strings"

	"github.com/zackslash/sansseed/derivation"
	"github.com/zackslash/sansseed/lengths"
	"github.com/zackslash/sansseed/wordlists"
)

// New24WordMnemonicPhraseForLanguage is a GoMobile compatible function that
// returns a new random 24 word mnemonic phrase for given language as a single space seperated string
func New24WordMnemonicPhraseForLanguage(language string) (string, error) {
	ent, err := NewWordEntropy(lengths.SeedBitLength(lengths.TwentyfourWordSeed))
	if err != nil {
		return "", err
	}

	wl, err := wordlists.GetByLanguageName(language)
	if err != nil {
		return "", err
	}

	reSlice, err := MnemonicPhraseForLanguage(ent, *wl)
	if err != nil {
		return "", err
	}

	r := strings.Join(reSlice, " ")
	return r, nil
}

// DeriveSeedFromMnemonic is a GoMobile compatible function that
// returns the hex encoded seed derived from the given mnemonic
func DeriveSeedFromMnemonic(mnemonic string) string {
	seed := derivation.DeriveSeedFromMnemonic(mnemonic, "")
	return hex.EncodeToString(seed)
}
