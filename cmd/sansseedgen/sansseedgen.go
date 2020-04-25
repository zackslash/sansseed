package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/zackslash/sansseed"
	"github.com/zackslash/sansseed/derivation"
	"github.com/zackslash/sansseed/lengths"
	"github.com/zackslash/sansseed/wordlists"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Mnemonic generation
	gen  = kingpin.Flag("generate", "generates a new mnemonic phrase").Default().Short('g').Bool()
	lang = kingpin.Flag("language", "the language to use for mnemonic phrase generation").Default("english").Short('c').String()
	bulk = kingpin.Flag("bulk", "specify the number of seeds to generate (used for bulk generation of mnemonics)").Default("1").Short('b').Int64()
	len  = kingpin.Flag("length", "mnemonic phrase length to generate").Default("12").Short('l').Int32()

	// Seed derivation from mnemonic
	mnemonic = kingpin.Flag("mnemonic", "supplying a mnemonic phrase will derive the hex encoded 512bit seed").Default("").Short('m').String()
	password = kingpin.Flag("password", "password for mnemonic phrase (used with mnemonic to derive the seed)").Default("").Short('p').String()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()
	if *gen {
		for index := int64(0); index < *bulk; index++ {
			generate(int(*len), *lang)
		}
		return
	}

	if *mnemonic != "" {
		seed := derivation.DeriveSeedFromMnemonic(*mnemonic, *password)
		fmt.Printf("%s\n", hex.EncodeToString(seed))
		return
	}
}

// generate a new mnemonic phrase
func generate(len int, language string) {
	l := []string{}
	for index := 0; index < len; index++ {
		l = append(l, "N")
	}
	p := strings.Join(l, " ")
	t, err := lengths.GetLengthTypeForMnemonic(p)
	if err != nil {
		fmt.Printf("Failed to generate: %s\n", err.Error())
		return
	}
	if t == lengths.UnknownWordSeed {
		fmt.Printf("Invalid seed length\n")
		return
	}

	ent, err := sansseed.NewWordEntropy(t)
	if err != nil {
		fmt.Printf("Failed to generate: %s\n", err.Error())
		return
	}

	wordList, err := wordlists.GetByLanguageName(language)
	if err != nil {
		fmt.Printf("Failed to find language '%s' check documentation for available languages\n", language)
		return
	}

	result, err := sansseed.MnemonicPhraseForLanguage(ent, *wordList)
	if err != nil {
		fmt.Printf("Failed to generate: %s\n", err.Error())
		return
	}

	res := strings.Join(result, " ")
	fmt.Printf("%s\n", res)
}
