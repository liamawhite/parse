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
	. "github.com/liamawhite/parse/test"
)

func TestWhitespace(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "Whitespace: no match",
			Input:          "ABCDEF",
			Parser:         Whitespace,
			ExpectedMatch:  "",
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "Whitespace: match",
			Input:          " \t \n \r ABC",
			Parser:         Whitespace,
			ExpectedMatch:  " \t \n \r ",
			ExpectedOK:     true,
			RemainingInput: "ABC",
		},
		{
			Name:           "Inlined whitespace: no match",
			Input:          "\nABC",
			Parser:         InlineWhitespace,
			ExpectedOK:     false,
			RemainingInput: "\nABC",
		},
		{
			Name:           "Inlined whitespace: match",
			Input:          " \t ABC",
			Parser:         InlineWhitespace,
			ExpectedMatch:  " \t ",
			ExpectedOK:     true,
			RemainingInput: "ABC",
		},
	}
	RunTests(t, tests)
}

func TestOptionalWhitespace(t *testing.T) {
	tests := []ParserTest[string]{
		{
			Name:           "OptionalWhitespace: no match",
			Input:          "ABCDEF",
			Parser:         OptionalWhitespace,
			ExpectedMatch:  "",
			ExpectedOK:     true,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "OptionalWhitespace: match",
			Input:          " 	ABC",
			Parser:         OptionalWhitespace,
			ExpectedMatch:  " 	",
			ExpectedOK:     true,
			RemainingInput: "ABC",
		},
		{
			Name:           "OptionalInlineWhitespace: no match",
			Input:          "\nABC",
			Parser:         OptionalInlineWhitespace,
			ExpectedMatch:  "",
			ExpectedOK:     true,
			RemainingInput: "\nABC",
		},
		{
			Name:           "OptionalInlineWhitespace: match",
			Input:          " \t ABC",
			Parser:         OptionalInlineWhitespace,
			ExpectedMatch:  " \t ",
			ExpectedOK:     true,
			RemainingInput: "ABC",
		},
	}
	RunTests(t, tests)
}
