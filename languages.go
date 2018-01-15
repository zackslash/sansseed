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

import "github.com/sanscentral/sansseed/wordlists/languages"

// MnemonicPhraseChineaseSimplified uses given entropy to
// return a mnemonic phrase for the Chinease Simplified language
func MnemonicPhraseChineaseSimplified(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39ChineaseSimplified{})
}

// MnemonicPhraseChineaseTraditional uses given entropy to
// return a mnemonic phrase for the Chinease Traditional language
func MnemonicPhraseChineaseTraditional(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39ChineaseTraditional{})
}

// MnemonicPhraseEnglish uses given entropy to
// return a mnemonic phrase for the English language
func MnemonicPhraseEnglish(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39English{})
}

// MnemonicPhraseFrench uses given entropy to
// return a mnemonic phrase for the French language
func MnemonicPhraseFrench(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39French{})
}

// MnemonicPhraseItalian uses given entropy to
// return a mnemonic phrase for the Italian language
func MnemonicPhraseItalian(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39Italian{})
}

// MnemonicPhraseJapanese uses given entropy to
// return a mnemonic phrase for the Japanese language
func MnemonicPhraseJapanese(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39Japanese{})
}

// MnemonicPhraseKorean uses given entropy to
// return a mnemonic phrase for the Korean language
func MnemonicPhraseKorean(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39Korean{})
}

// MnemonicPhraseSpanish uses given entropy to
// return a mnemonic phrase for the Spanish language
func MnemonicPhraseSpanish(i []int, e error) ([]string, error) {
	if e != nil {
		return nil, e
	}
	return MnemonicPhraseForLanguage(i, languages.BIP39Spanish{})
}
