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

func times[T any](min int, max func(i int) bool, p Parser[T]) Parser[[]T] {
	return func(in Input) ([]T, bool, error) {
		start := in.Checkpoint()
		match := make([]T, 0)
		for i := 0; max(i); i++ {
			m, ok, err := p(in)
			if err != nil {
				in.Restore(start)
				return match, false, err
			}
			if !ok {
				break
			}
			match = append(match, m)
		}
		ok := len(match) >= min && max(len(match)-1)
		if !ok {
			in.Restore(start)
			return nil, false, nil
		}
		return match, true, nil
	}
}

// Matches the given parser exactly n times.
func Times[T any](n int, p Parser[T]) Parser[[]T] {
	return times(n, func(i int) bool { return i < n }, p)
}

// Matches the given parser between min and max times .
func Between[T any](min, max int, p Parser[T]) Parser[[]T] {
	return times(min, func(i int) bool { return i < max }, p)
}

// Matches the given parser at least n times.
func AtLeast[T any](n int, p Parser[T]) Parser[[]T] {
	return times(n, func(i int) bool { return true }, p)
}

// Matches the given parser at most n times.
func AtMost[T any](n int, p Parser[T]) Parser[[]T] {
	return times(0, func(i int) bool { return i < n }, p)
}

// Matches the given parser zero or more times.
func ZeroOrMore[T any](p Parser[T]) Parser[[]T] {
	return times(0, func(i int) bool { return true }, p)
}

// Matches the given parser one or more times.
func OneOrMore[T any](p Parser[T]) Parser[[]T] {
	return times(1, func(i int) bool { return true }, p)
}
