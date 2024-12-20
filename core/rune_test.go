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
	"unicode"

	"github.com/liamawhite/parse/core"
	. "github.com/liamawhite/parse/test"
)

func TestRuneWhere(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "RuneWhere: no match",
			Input:          "ABCDEF",
			Parser:         core.RuneWhere(func(r rune) bool { return r == 'a' }),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:  "RuneWhere: match",
			Input: "ABCDEF",
			Parser: core.RuneWhere(func(r rune) bool {
				return unicode.IsUpper(r)
			}),
			ExpectedMatch:  "A",
			ExpectedOK:     true,
			RemainingInput: "BCDEF",
		},
		{
			Name:           "AnyRune: match",
			Input:          "ABCDEF",
			Parser:         core.AnyRune,
			ExpectedMatch:  "A",
			ExpectedOK:     true,
			RemainingInput: "BCDEF",
		},
		{
			Name:           "RuneIn: no match",
			Input:          "ABCDEF",
			Parser:         core.RuneIn("123"),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "RuneIn: match",
			Input:          "ABCDEF",
			Parser:         core.RuneIn("CBA"),
			ExpectedMatch:  "A",
			ExpectedOK:     true,
			RemainingInput: "BCDEF",
		},
		{
			Name:           "RuneNotIn: no match",
			Input:          "ABCDEF",
			Parser:         core.RuneNotIn("ABC"),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "RuneNotIn: match",
			Input:          "ABCDEF",
			Parser:         core.RuneNotIn("123"),
			ExpectedMatch:  "A",
			ExpectedOK:     true,
			RemainingInput: "BCDEF",
		},
		{
			Name:           "RuneInRanges: match",
			Input:          "     ",
			Parser:         core.RuneInRanges(unicode.White_Space),
			ExpectedMatch:  " ",
			ExpectedOK:     true,
			RemainingInput: "    ",
		},
		{
			Name:           "RuneInRanges: no match",
			Input:          "     ",
			Parser:         core.RuneInRanges(unicode.Han),
			ExpectedOK:     false,
			RemainingInput: "     ",
		},
		{
			Name:          "Letter: match",
			Input:         "a",
			Parser:        core.Letter,
			ExpectedMatch: "a",
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestDigit(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "Digit: no match",
			Input:          "ABCDEF",
			Parser:         core.Digit,
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "Digit: nine",
			Input:          "987",
			Parser:         core.Digit,
			ExpectedMatch:  "9",
			ExpectedOK:     true,
			RemainingInput: "87",
		},
		{
			Name:           "Digit: zero",
			Input:          "0123",
			Parser:         core.Digit,
			ExpectedMatch:  "0",
			ExpectedOK:     true,
			RemainingInput: "123",
		},
	}
	RunTests(t, tests)
}
