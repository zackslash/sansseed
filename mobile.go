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

package sansseed

import (
	"encoding/hex"
	"strings"

	"github.com/sanscentral/sansseed/derivation"
	"github.com/sanscentral/sansseed/lengths"
	"github.com/sanscentral/sansseed/wordlists"
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
