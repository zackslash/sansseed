package lengths_test

import (
	"testing"

	"github.com/zackslash/sansseed/lengths"
)

func TestInvalidMnemonicLength(t *testing.T) {
	r, err := lengths.GetLengthTypeForMnemonic("ONE TWO THREE")
	if err == nil {
		t.Error("No error when expected for invalid mnemonic phrase length")
	}
	if r != lengths.UnknownWordSeed {
		t.Error("Incorrect type for invalid mnemonic phrase length")
	}
}
