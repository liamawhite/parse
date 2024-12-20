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

func TestUntil(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			Name:           "Until: success",
			Input:          "ABCDEF",
			Parser:         core.Until(core.AnyRune, core.Rune('D')),
			ExpectedMatch:  []string{"A", "B", "C"},
			ExpectedOK:     true,
			RemainingInput: "DEF",
		},
		{
			Name:           "Until: reach end before delimiter",
			Input:          "ABC",
			Parser:         core.Until(core.AnyRune, core.Rune('D')),
			ExpectedOK:     false,
			RemainingInput: "ABC",
		},
		{
			Name:           "Until: delimiter matches on first character",
			Input:          "ABCA",
			Parser:         core.Until(core.AnyRune, core.Rune('A')),
			ExpectedMatch:  []string{"A", "B", "C"},
			ExpectedOK:     true,
			RemainingInput: "A",
		},
		{
			Name:          "UntilEOF: allow EOF",
			Input:         "ABC",
			Parser:        core.UntilEOF(core.AnyRune, core.Digit),
			ExpectedMatch: []string{"A", "B", "C"},
			ExpectedOK:    true,
		},
		{
			Name:           "UntilEOF: delimiter reached",
			Input:          "ABC1",
			Parser:         core.UntilEOF(core.AnyRune, core.Digit),
			ExpectedOK:     true,
			ExpectedMatch:  []string{"A", "B", "C"},
			RemainingInput: "1",
		},
		{
			Name:       "UntilEOF: empty input",
			Input:      "",
			Parser:     core.UntilEOF(core.AnyRune, core.Digit),
			ExpectedOK: false, // needs at least one match before EOF
		},
	}
	RunTests(t, tests)

}
