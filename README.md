# SansSeed

SansSeed is both a library and command-line tool for generating BIP39 compatible mnemonic phrases and derivation of BIP39 seeds for research purposes.

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
