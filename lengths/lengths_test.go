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

package lengths_test

import (
	"testing"

	"github.com/sanscentral/sansseed/lengths"
)

func TestInvalidMnemonicLength(t *testing.T) {
	r, err := lengths.GetLengthTypeForMnemonic("ONE TWO THREE")
	if err == nil {
		t.Error("No error when expected for invalid mnemonic phrase length")
	}
	if r != lengths.UnknownWordSeed {
		t.Error("Incorrect type for invalid mnemonic phrase length")
	}
}
