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

package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/sanscentral/sansseed"
	"github.com/sanscentral/sansseed/derivation"
	"github.com/sanscentral/sansseed/lengths"
	"github.com/sanscentral/sansseed/wordlists"

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
