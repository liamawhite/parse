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

func TestOptional(t *testing.T) {
	tests := []ParserTest[core.Match[string]]{
		{
			Name:           "Optional: it's not there, but that's OK",
			Input:          "ABCDEF",
			Parser:         core.Optional(core.String("1")),
			ExpectedMatch:  core.NewMatch("", false),
			ExpectedOK:     true,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "Optional: it's there, so return the value",
			Input:          "ABCDEF",
			Parser:         core.Optional(core.String("A")),
			ExpectedMatch:  core.NewMatch("A", true),
			ExpectedOK:     true,
			RemainingInput: "BCDEF",
		},
	}
	RunTests(t, tests)
}
