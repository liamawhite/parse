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

type Match[T any] interface {
	Values() T
	Ok() bool
}

type match[T any] struct {
	value T
	ok    bool
}

func NewMatch[T any](value T, ok bool) Match[T] {
	return match[T]{value: value, ok: ok}
}

func (m match[T]) Values() T {
	return m.value
}

func (m match[T]) Ok() bool {
	return m.ok
}

// Optional wraps a parser, returning a match struct with Ok set to true if the parser matches, otherwise Ok is set to false.
// The top level parser will always return true, unless an error occurs.
func Optional[T any](parser Parser[T]) Parser[Match[T]] {
	return func(in Input) (Match[T], bool, error) {
		m, ok, err := parser(in)
		if err != nil {
			return match[T]{}, false, err
		}
		return match[T]{value: m, ok: ok}, true, nil
	}
}
