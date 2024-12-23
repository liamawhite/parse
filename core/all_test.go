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

func TestAll(t *testing.T) {
	tests := []ParserTest[[]string]{
		{
			Name:           "no match",
			Input:          "AC",
			Parser:         core.All(core.Rune('A'), core.Rune('B')),
			ExpectedOK:     false,
			RemainingInput: "AC",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "AC",
			Parser:         core.All(NaughtyParser[string](), core.Rune('C')),
			ExpectedOK:     false,
			RemainingInput: "AC",
		},
		{
			Name:          "match",
			Input:         "AB",
			Parser:        core.All(core.Rune('A'), core.Rune('B')),
			ExpectedMatch: []string{"A", "B"},
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
