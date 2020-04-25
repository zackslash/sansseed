# SansSeed

SansSeed is both a library and command-line tool for generating BIP39 compatible mnemonic phrases and derivation of BIP39 seeds.

## Usage

#### CLI

See [Here](https://github.com/zackslash/sansseed/tree/master/cmd/sansseedgen) for details on how to use the `sansseedgen` CLI tool.

#### Library

##### Generating a mnemonic

```go
    // Generate new random entropy for a 24 word seed
    ent, err := sansseed.NewWordEntropy(lengths.SeedBitLength(lengths.TwentyfourWordSeed))
    if err != nil {
        return err
    }

    // Generate an english mnemonic using the new entropy (Note this package currently supports 8 languages)
    res, err := sansseed.MnemonicPhraseForLanguage(ent, languages.BIP39English{})
    if err != nil {
        return err
    }
```

##### Derive seed from mnemonic

```go
    optionalPassword := "S3CRET"
    mnemonic := "vessel ladder alter error federal sibling chat ability sun glass valve picture"
    seed := derivation.DeriveSeedFromMnemonic(mnemonic, optionalPassword)
```

## Build

Requires Go version 1.13 or later.

Run all unit tests with `$ go test ./...`

## License

The MIT License

Copyright (c) 2020 Luke Hines

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
