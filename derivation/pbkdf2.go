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

package derivation

import (
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

var (
	salt       = "mnemonic"
	iterations = 2048
	keyLength  = 64
)

// DeriveSeedFromMnemonic turns a mnemonic into a 512bit seed
func DeriveSeedFromMnemonic(mnemonic, optionalPassword string) []byte {
	s := salt + optionalPassword
	return pbkdf2.Key([]byte(mnemonic), []byte(s), iterations, keyLength, sha512.New)
}
