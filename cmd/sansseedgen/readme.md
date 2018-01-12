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

## Contact

contact@sanscentral.org ([PGP](../../resources/publickey.contact@sanscentral.org.asc))

## License

![AGPLv3 Logo](../../resources/agplv3-155x51.png)

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
