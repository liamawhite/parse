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

import "unicode"

// Whitespace parses whitespace.
var Whitespace Parser[string] = StringFrom(OneOrMore(RuneInRanges(unicode.White_Space)))

// InlineWhitespace parses inline whitespace (spaces and tabs).
var InlineWhitespace = StringFrom(OneOrMore(RuneIn(" \t")))

// OptionalWhitespace parses optional whitespace.
var OptionalWhitespace = func(in Input) (output string, ok bool, err error) {
	output, ok, err = Whitespace(in)
	if err != nil {
		return
	}
	return output, true, nil
}

// OptionalInlineWhitespace parses optional inline whitespace.
var OptionalInlineWhitespace = func(in Input) (output string, ok bool, err error) {
	output, ok, err = InlineWhitespace(in)
	if err != nil {
		return
	}
	return output, true, nil
}

// CR is a carriage return.
var CR = Rune('\r')

// CR parses a line feed, used by Unix systems as the newline.
var LF = Rune('\n')

// CRLF parses a carriage returned, followed by a line feed, used by Windows systems as the newline.
var CRLF = String("\r\n")

// NewLine matches either a Windows or Unix line break character.
var NewLine = Any(CRLF, LF)
