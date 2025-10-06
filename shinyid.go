// Copyright 2025 itpey
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shinyid

import "errors"

const base64URLCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

var (
	decodeIdx [256]byte
)

func init() {
	// initialize all to 0xFF (invalid)
	for i := range decodeIdx {
		decodeIdx[i] = 0xFF
	}
	for i := 0; i < len(base64URLCharset); i++ {
		decodeIdx[base64URLCharset[i]] = byte(i)
	}
}

// ToShiny converts an id to a shiny.
func ToShiny(id uint64) string {
	// Max length: ceil(64/6)=11
	var b [11]byte
	i := len(b) - 1
	for {
		b[i] = base64URLCharset[id&0x3F]
		id >>= 6
		if id == 0 {
			return string(b[i:])
		}
		i--
	}
}

// ToId converts a shiny to its corresponding id.
func ToId(shiny string) (uint64, error) {
	if len(shiny) == 0 || len(shiny) > 11 {
		return 0, errors.New("invalid shiny")
	}
	var v uint64
	for i := 0; i < len(shiny); i++ {
		c := shiny[i]
		idx := decodeIdx[c]
		if idx == 0xFF {
			return 0, errors.New("invalid shiny")
		}
		v = (v << 6) | uint64(idx)
	}
	return v, nil
}
