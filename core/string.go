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

import "strings"

// String matches the given string (case sensitive).
func String(s string) Parser[string] {
	return stringWhere(s, func(candidate string) bool {
		return s == candidate
	})
}

// StringInsensitive matches the given string (case insensitive).
func StringInsensitive(s string) Parser[string] {
	return stringWhere(s, func(candidate string) bool {
		return strings.EqualFold(s, candidate)
	})
}

func stringWhere(s string, predicate func(candidate string) bool) Parser[string] {
	return func(in Input) (string, bool, error) {
		match, ok := in.Peek(len(s))
		if !ok {
			return "", false, nil
		}
		if !predicate(match) {
			return "", false, nil
		}
		res, _ := in.Take(len(s))
		return res, true, nil
	}
}

// StringFrom returns the string range match by the given parsers.
func StringFrom[T any](parsers ...Parser[T]) Parser[string] {
	return func(in Input) (string, bool, error) {
		start := in.Checkpoint()
		for _, parser := range parsers {
			_, ok, err := parser(in)
			if err != nil {
				return "", false, err
			}
			if !ok {
				in.Restore(start)
				return "", false, nil
			}
		}

		end := in.Checkpoint()
		in.Restore(start)
		res, ok := in.Take(int(end - start))
		return res, ok, nil
	}
}

// StringUntil matches until the delimiter is found.
// This differs from while as the delimiter is searched for AFTER the first match.
func StringUntil[T any](delimiter Parser[T]) Parser[string] {
	return func(in Input) (string, bool, error) {
		start := in.Checkpoint()
		for {
			_, chompOk := in.Take(1)
			if !chompOk {
				in.Restore(start)
				return "", false, nil
			}

			beforeDelimiter := in.Checkpoint()
			_, ok, err := delimiter(in)
			if err != nil {
				in.Restore(start)
				return "", false, err
			}
			if ok {
				in.Restore(beforeDelimiter)
				break
			}
		}
		end := in.Checkpoint()
		in.Restore(start)
		res, ok := in.Take(int(end - start))
		return res, ok, nil
	}
}

// StringUntilEOF matches until the delimiter is found or the end of the input.
// This requires at least one character to be matched before the end of the input.
// See the tests for examples.
func StringUntilEOF[T any](delimiter Parser[T]) Parser[string] {
	return func(in Input) (string, bool, error) {
		return StringUntil(Or(delimiter, EOF[string]()))(in)
	}
}

// StringWhileNot matches while the delimiter does not match.
// This differs from until as the delimiter is searched for BEFORE the first match.
func StringWhileNot[T any](delimiter Parser[T]) Parser[string] {
	return func(in Input) (string, bool, error) {
		start := in.Checkpoint()
		for {
			beforeDelimiter := in.Checkpoint()
			_, ok, err := delimiter(in)
			if err != nil {
				in.Restore(start)
				return "", false, err
			}
			if ok {
				in.Restore(beforeDelimiter)
				break
			}

			_, chompOk := in.Take(1)
			if !chompOk {
				in.Restore(start)
				return "", false, nil
			}
		}
		end := in.Checkpoint()
		in.Restore(start)
		res, ok := in.Take(int(end - start))
		return res, ok, nil
	}
}

// StringWhileNotOrEOF matches while the delimiter does not match or the end of the input.
// This does not require any characters to be matched before the end of the input.
func StringWhileNotEOFOr[T any](delimiter Parser[T]) Parser[string] {
	return func(in Input) (string, bool, error) {
		return StringWhileNot(Or(delimiter, EOF[string]()))(in)
	}
}
