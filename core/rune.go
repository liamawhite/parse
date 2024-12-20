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

package core

import (
	"strings"
	"unicode"
)

// Rune matches a single rune.
func Rune(r rune) Parser[string] {
	return RuneWhere(func(candidate rune) bool {
		return candidate == r
	})
}

// RuneWhere matches a single rune when the predicate is true.
func RuneWhere(predicate func(r rune) bool) Parser[string] {
	return func(in Input) (string, bool, error) {
		match, ok := in.Peek(1)
		if !ok {
			return "", false, nil
		}
		if !predicate(rune(match[0])) {
			return "", false, nil
		}
		res, _ := in.Take(1)
		return res, true, nil
	}
}

// RuneIn matches a single rune when the rune is in the given string.
func RuneIn(s string) Parser[string] {
	return RuneWhere(func(r rune) bool {
		return strings.Contains(s, string(r))
	})
}

// RuneNotIn matches a single rune when the rune is not in the given string.
func RuneNotIn(s string) Parser[string] {
	return RuneWhere(func(r rune) bool {
		return !strings.Contains(s, string(r))
	})
}

// RuneInRanges matches a single rune when the rune is in one of the given unicode ranges.
func RuneInRanges(ranges ...*unicode.RangeTable) Parser[string] {
	return RuneWhere(func(r rune) bool { return unicode.IsOneOf(ranges, r) })
}

// AnyRune matches any single rune.
var AnyRune = RuneWhere(func(r rune) bool { return true })

// Letter matches any rune within the letter unicode range.
var Letter = RuneInRanges(unicode.Letter)

// Digit matches any rune within the number unicode range.
var Digit = RuneInRanges(unicode.Number)
