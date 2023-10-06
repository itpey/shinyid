// Copyright 2023 'itpey'
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

import (
	"encoding/base64"
	"testing"
)

func TestToShiny(t *testing.T) {
	tests := []struct {
		id           uint64
		expectedCode string
	}{
		{0, "A"},
		{62, "-"},
		{63, "_"},
		{1, "B"},
		{500, "H0"},
		{9375, "CSf"},
		{18446744073709551615, "P__________"},
	}

	for _, test := range tests {
		result := ToShiny(test.id)
		if result != test.expectedCode {
			t.Errorf("ToShiny(%d) = %s, expected %s", test.id, result, test.expectedCode)
		}
	}
}

func TestToID(t *testing.T) {
	tests := []struct {
		code          string
		expectedId    uint64
		expectedError bool
	}{
		{"A", 0, false},
		{"-", 62, false},
		{"_", 63, false},
		{"B", 1, false},
		{"H0", 500, false},
		{"CSf", 9375, false},
		{"P__________", 18446744073709551615, false},
		{"#", 100, true},
		{"1_$_", 200, true},
	}

	for _, test := range tests {
		result, err := ToId(test.code)
		if err != nil {
			if !test.expectedError {
				t.Errorf("ToId(%s) returned an unexpected error: %v", test.code, err)
			}
			continue
		}

		if result != test.expectedId {
			t.Errorf("ToId(%s) = %d, expected %d", test.code, result, test.expectedId)
		}
	}
}

func TestIsValidShiny(t *testing.T) {
	tests := []struct {
		code         string
		expectedBool bool
	}{
		{"A", true},
		{"b", true},
		{"_", true},
		{"-", true},
		{"F_0", true},
		{"fPg97-", true},
		{"!@#$%", false},
		{"1_1-Q_@-", false},
	}

	for _, test := range tests {
		result := isValidShiny(test.code)
		if result != test.expectedBool {
			t.Errorf("isValidShiny(%s) = %v, expected %v", test.code, result, test.expectedBool)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		char         rune
		expectedBool bool
	}{
		{'A', true},
		{'a', true},
		{'Z', true},
		{'z', true},
		{'0', true},
		{'9', true},
		{'#', false},
		{'@', false},
		{'$', false},
	}

	for _, test := range tests {
		result := isAlphaNumeric(test.char)
		if result != test.expectedBool {
			t.Errorf("isAlphaNumeric(%c) = %v, expected %v", test.char, result, test.expectedBool)
		}
	}
}

func BenchmarkShinyIDEncoding(b *testing.B) {
	id := uint64(18446744073709551615)
	for i := 0; i < b.N; i++ {
		ToShiny(id)
	}
}

func BenchmarkShinyIDDecoding(b *testing.B) {
	shiny := "P__________"
	for i := 0; i < b.N; i++ {
		ToId(shiny)
	}
}

func BenchmarkBase64Encoding(b *testing.B) {
	data := []byte("18446744073709551615")
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString(data)
	}
}

func BenchmarkBase64Decoding(b *testing.B) {
	data := "MTg0NDY3NDQwNzM3MDk1NTE2MTU="
	for i := 0; i < b.N; i++ {
		_, _ = base64.StdEncoding.DecodeString(data)
	}
}
