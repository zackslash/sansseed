package sansseed

import (
	"errors"

	"github.com/zackslash/sansseed/entropy"
	"github.com/zackslash/sansseed/lengths"
	"github.com/zackslash/sansseed/wordlists"
)

// NewWordEntropy returns new mnemonic integer slice of given length
func NewWordEntropy(l lengths.SeedBitLength) ([]int, error) {
	e, err := entropy.GetRandomEntropyBytesWithCheckSum(l)
	if err != nil {
		return []int{}, err
	}

	ints, err := entropy.BinaryStringToIntSlice(e)
	if err != nil {
		return []int{}, err
	}

	return ints, nil
}

// MnemonicPhraseForLanguage uses given entropy and language to generate a mnemonic phrase
func MnemonicPhraseForLanguage(i []int, l wordlists.BIP39Wordlist) ([]string, error) {
	w := l.GetList()
	r := []string{}
	for _, en := range i {
		if en <= 0 || en >= len(w) {
			return []string{}, errors.New("entropy exceeds bounds of word list")
		}

		wrd := w[en]
		r = append(r, wrd)
	}
	return r, nil
}

// New12WordEntropy returns new twelve length seed as an integer slice
func New12WordEntropy() ([]int, error) {
	return NewWordEntropy(lengths.TwelveWordSeed)
}

// New15WordEntropy returns new fiveteen length seed as an integer slice
func New15WordEntropy() ([]int, error) {
	return NewWordEntropy(lengths.FiveteenWordSeed)
}

// New18WordEntropy returns new eighteen length seed as an integer slice
func New18WordEntropy() ([]int, error) {
	return NewWordEntropy(lengths.EighteenWordSeed)
}

// New21WordEntropy returns new twentyone length seed as an integer slice
func New21WordEntropy() ([]int, error) {
	return NewWordEntropy(lengths.TwentyoneWordSeed)
}

// New24WordEntropy returns new twentyfour length seed as an integer slice
func New24WordEntropy() ([]int, error) {
	return NewWordEntropy(lengths.TwentyfourWordSeed)
}
