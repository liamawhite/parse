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

	"github.com/liamawhite/parse/core"
	. "github.com/liamawhite/parse/test"
)

func TestString(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "no match",
			Input:          "ABCDEF",
			Parser:         core.String("123"),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "matches",
			Input:          "ABCDEF",
			Parser:         core.String("ABC"),
			ExpectedMatch:  "ABC",
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
		{
			Name:           "matches insensitive",
			Input:          "ABCDEF",
			Parser:         core.StringInsensitive("abc"),
			ExpectedMatch:  "ABC",
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
		{
			Name:           "matches insensitive (inverse)",
			Input:          "abCDEF",
			Parser:         core.StringInsensitive("ABC"),
			ExpectedMatch:  "abC",
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
	}
	RunTests(t, tests)
}

func TestStringFrom(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "no match",
			Input:          "ABCDEF",
			Parser:         core.StringFrom(core.String("ABC"), core.String("123")),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:          "matches",
			Input:         "ABCDEF",
			Parser:        core.StringFrom(core.String("A"), core.String("BC"), core.String("DEF")),
			ExpectedMatch: "ABCDEF",
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestStringUntil(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "StringUntil: success",
			Input:          "ABCDEF",
			Parser:         core.StringUntil(core.String("D")),
			ExpectedMatch:  "ABC",
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
		{
			Name:           "StringUntil: fail, reached EOF before delimiter was found",
			Input:          "ABCDEF",
			Parser:         core.StringUntil(core.String("G")),
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "StringUntil: success, delimiter matches first character",
			Input:          "ABCDEFA",
			Parser:         core.StringUntil(core.String("A")),
			ExpectedMatch:  "ABCDEF",
			ExpectedOK:     true,
			RemainingInput: "A",
		},
		{
			Name:           "StringUntilEOF: stop at the delimiter if it's there",
			Input:          "ABCDEF",
			Parser:         core.StringUntilEOF(core.String("F")),
			ExpectedMatch:  "ABCDE",
			ExpectedOK:     true,
			RemainingInput: "F",
		},
		{
			Name:          "StringUntilEOF: allow EOF",
			Input:         "ABCDEF",
			Parser:        core.StringUntilEOF(core.String("G")),
			ExpectedMatch: "ABCDEF",
			ExpectedOK:    true,
		},
		{
			Name:          "StringUntilEOF: allow EOF, empty input",
			Input:         "",
			Parser:        core.StringUntilEOF(core.String("G")),
			ExpectedMatch: "", // expects a character to be matched before EOF
			ExpectedOK:    false,
		},
	}
	RunTests(t, tests)
}

func TestStringWhileNot(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "StringWhileNot: success",
			Input:          "ABCDEF1",
			Parser:         core.StringWhileNot(core.Digit),
			ExpectedMatch:  "ABCDEF",
			ExpectedOK:     true,
			RemainingInput: "1",
		},
		{
			Name:           "StringWhileNot: fail, reached EOF before delimiter was found",
			Input:          "ADBCDEF",
			Parser:         core.StringWhileNot(core.Digit),
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ADBCDEF",
		},
		{
			Name:           "StringWhileNot: success, delimiter matches first character",
			Input:          "1ABCDEF",
			Parser:         core.StringWhileNot(core.Digit),
			ExpectedMatch:  "",
			ExpectedOK:     true,
			RemainingInput: "1ABCDEF",
		},
		{
			Name:           "StringWhileNotEOF: stop at the delimiter if it's there",
			Input:          "ABCDEF1",
			Parser:         core.StringWhileNotEOFOr(core.Digit),
			ExpectedMatch:  "ABCDEF",
			ExpectedOK:     true,
			RemainingInput: "1",
		},
		{
			Name:           "StringWhileNotEOF: allow EOF",
			Input:          "ABCDEF",
			Parser:         core.StringWhileNotEOFOr(core.Digit),
			ExpectedMatch:  "ABCDEF",
			ExpectedOK:     true,
			RemainingInput: "",
		},
		{
			Name:          "StringWhileNotEOF: allow EOF, empty input",
			Input:         "",
			Parser:        core.StringWhileNotEOFOr(core.Digit),
			ExpectedMatch: "",
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
