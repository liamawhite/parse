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

func TestOr(t *testing.T) {
	tests := []ParserTest[Tuple2[Match[string], Match[string]]]{
		{
			Name:           "no match",
			Input:          "C",
			Parser:         Or(Rune('A'), Rune('B')),
			ExpectedMatch:  NewTuple2(NewMatch("", false), NewMatch("", false)),
			ExpectedOK:     false,
			RemainingInput: "C",
		},
		{
			Name:          "first match",
			Input:         "A",
			Parser:        Or(Rune('A'), Rune('B')),
			ExpectedMatch: NewTuple2[Match[string], Match[string]](NewMatch("A", true), nil),
			ExpectedOK:    true,
		},
		{
			Name:          "second match",
			Input:         "B",
			Parser:        Or(Rune('A'), Rune('B')),
			ExpectedMatch: NewTuple2(NewMatch("", false), NewMatch("B", true)),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
