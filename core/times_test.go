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

func TestTime(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			Name:           "Times: no match",
			Input:          "ABCDEF",
			Parser:         core.Times(2, core.Rune('A')),
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:          "Times: match",
			Input:         "AA",
			Parser:        core.Times(2, core.Rune('A')),
			ExpectedMatch: []string{"A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:           "Times: no match rolls back input even if one of the parsers consumed input",
			Input:          "AC",
			Parser:         core.Times(2, NaughtyParser[string]()),
			ExpectedOK:     false,
			RemainingInput: "AC",
		},
		{
			Name:          "Between: at least 1 up to 5 with 4",
			Input:         "AAAA",
			Parser:        core.Between(1, 5, core.Rune('A')),
			ExpectedMatch: []string{"A", "A", "A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:          "Between: at least 1 up to 5 with 5",
			Input:         "AAAAA",
			Parser:        core.Between(1, 5, core.Rune('A')),
			ExpectedMatch: []string{"A", "A", "A", "A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:           "Between: at least 1 up to 5 with 6",
			Input:          "AAAAAA",
			Parser:         core.Between(1, 5, core.Rune('A')),
			ExpectedMatch:  []string{"A", "A", "A", "A", "A"},
			ExpectedOK:     true,
			RemainingInput: "A",
		},
		{
			Name:       "Between: at least 1 up to 5 with 0",
			Input:      "",
			Parser:     core.Between(1, 5, core.Rune('A')),
			ExpectedOK: false,
		},
		{
			Name:          "Between: at least 1 up to 5 with 1",
			Input:         "A",
			Parser:        core.Between(1, 5, core.Rune('A')),
			ExpectedMatch: []string{"A"},
			ExpectedOK:    true,
		},
		{
			Name:       "AtLeast: at least 1 with 0",
			Input:      "",
			Parser:     core.AtLeast(1, core.Rune('A')),
			ExpectedOK: false,
		},
		{
			Name:          "AtLeast: at least 1 with 1",
			Input:         "A",
			Parser:        core.AtLeast(1, core.Rune('A')),
			ExpectedMatch: []string{"A"},
			ExpectedOK:    true,
		},
		{
			Name:          "AtLeast: at least 1 with 2",
			Input:         "AA",
			Parser:        core.AtLeast(1, core.Rune('A')),
			ExpectedMatch: []string{"A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:          "AtMost: at most 3 with 0",
			Input:         "",
			Parser:        core.AtMost(3, core.Rune('A')),
			ExpectedMatch: []string{},
			ExpectedOK:    true,
		},
		{
			Name:          "AtMost: at most 3 with 1",
			Input:         "A",
			Parser:        core.AtMost(3, core.Rune('A')),
			ExpectedMatch: []string{"A"},
			ExpectedOK:    true,
		},
		{
			Name:          "AtMost: at most 3 with 3",
			Input:         "AAA",
			Parser:        core.AtMost(3, core.Rune('A')),
			ExpectedMatch: []string{"A", "A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:           "AtMost: at most 3 with 4",
			Input:          "AAAA",
			Parser:         core.AtMost(3, core.Rune('A')),
			ExpectedMatch:  []string{"A", "A", "A"},
			ExpectedOK:     true,
			RemainingInput: "A",
		},
		{
			Name:           "ZeroOrMore: no match",
			Input:          "BCDEF",
			Parser:         core.ZeroOrMore(core.Rune('A')),
			ExpectedOK:     true,
			ExpectedMatch:  []string{},
			RemainingInput: "BCDEF",
		},
		{
			Name:          "ZeroOrMore: match",
			Input:         "AA",
			Parser:        core.ZeroOrMore(core.Rune('A')),
			ExpectedMatch: []string{"A", "A"},
			ExpectedOK:    true,
		},
		{
			Name:           "OneOrMore: no match",
			Input:          "BCDEF",
			Parser:         core.OneOrMore(core.Rune('A')),
			ExpectedOK:     false,
			RemainingInput: "BCDEF",
		},
		{
			Name:          "OneOrMore: match",
			Input:         "AA",
			Parser:        core.OneOrMore(core.Rune('A')),
			ExpectedMatch: []string{"A", "A"},
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
