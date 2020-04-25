package sansseed

import "github.com/zackslash/sansseed/wordlists/languages"

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
