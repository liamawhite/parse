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

func TestWhileNot(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			Name:           "WhileNot: success",
			Input:          "ABCDEF",
			Parser:         core.WhileNot(core.AnyRune, core.Rune('D')),
			ExpectedMatch:  []string{"A", "B", "C"},
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
		{
			Name:           "WhileNot: reach end before delimiter",
			Input:          "ABC",
			Parser:         core.WhileNot(core.AnyRune, core.Rune('D')),
			ExpectedOK:     false,
			RemainingInput: "ABC",
		},
		{
			Name:           "WhileNot: delimiter matches on first character",
			Input:          "ABCA",
			Parser:         core.WhileNot(core.AnyRune, core.Rune('A')),
			ExpectedMatch:  []string{},
			ExpectedOK:     true,
			RemainingInput: "ABCA",
		},
		{
			Name:          "WhileNotEOF: allow EOF",
			Input:         "ABC",
			Parser:        core.WhileNotEOFOr(core.AnyRune, core.Rune('D')),
			ExpectedMatch: []string{"A", "B", "C"},
			ExpectedOK:    true,
		},
		{
			Name:           "WhileNotEOF: delimiter reached",
			Input:          "ABCD",
			Parser:         core.WhileNotEOFOr(core.AnyRune, core.Rune('D')),
			ExpectedMatch:  []string{"A", "B", "C"},
			ExpectedOK:     true,
			RemainingInput: "D",
		},
		{
			Name:          "WhileNotEOF: empty input",
			Input:         "",
			Parser:        core.WhileNotEOFOr(core.AnyRune, core.Rune('D')),
			ExpectedMatch: []string{},
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)

}
