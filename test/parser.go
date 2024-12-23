// Copyright 2024 Liam White
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

package core_test

import (
	"testing"

	. "github.com/liamawhite/parse/core"
	"github.com/stretchr/testify/assert"
)

type ParserTest[T any] struct {
	Name           string
	Input          string
	Parser         Parser[T]
	ExpectedMatch  T
	ExpectedOK     bool
	WantErr        bool
	RemainingInput string
}

func RunTests[T any](t *testing.T, tests []ParserTest[T]) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			in := NewInput(test.Input)
			match, ok, err := test.Parser(in)
			assert.Equal(t, test.ExpectedOK, ok)
			assert.Equal(t, test.ExpectedMatch, match)
			if test.WantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			remaining, _, _ := StringWhileNot(EOF[string]())(in)
			assert.Equal(t, test.RemainingInput, remaining)
		})
	}
}

// Change some of the characters in the string to uppercase
func CaPiTaLiZe(s string) string {
	runes := []rune(s)
	for i, r := range s {
		if i%2 == 0 && r >= 'a' && r <= 'z' {
			runes[i] = r - 32
		}
	}
	return string(runes)
}

// NaughtyParser parser consumes input but does not rollback
// This is useful for testing that combinators rollback input correctly
func NaughtyParser[T any]() Parser[T] {
	return func(in Input) (T, bool, error) {
		in.Take(1)
		var res T
		return res, false, nil
	}
}
