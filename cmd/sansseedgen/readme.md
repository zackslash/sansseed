# SansSeedGen

A command-line tool for generating BIP39 compatible mnemonic phrases and derivation of seeds

### Usage
```

usage: sansseedgen [<flags>]

Flags:
      --help                Show context-sensitive help (also try --help-long and --help-man).
  -g, --generate            generates a new mnemonic phrase
  -c, --language="english"  the language to use for mnemonic phrase generation
  -b, --bulk=1              specify the number of seeds to generate (used for bulk generation of mnemonics)
  -l, --length=12           mnemonic phrase length to generate
  -m, --mnemonic=""         supplying a mnemonic phrase will return the hex encoded 512bit seed
  -p, --password=""         password for mnemonic phrase (used with mnemonic to derive the seed)
      --version             Show application version.


Example: Generate 1 mnemonic
$ ./sansseedgen -g

Example: Generate 15 mnemonics
$ ./sansseedgen -g -b 15

Example: Derive seed from mnemonic
$ ./sansseedgen -m "army van defense carry jealous true garbage claim echo media make crunch"

Example: Derive seed from mnemonic with password
$ ./sansseedgen -m "army van defense carry jealous true garbage claim echo media make crunch" -p SuperDuperSecret
```

## License

