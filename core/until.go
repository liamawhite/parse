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

// Until matches until the delmiter is found.
// This differs from while as the delimiter is searched for AFTER the first match.
func Until[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return func(in Input) ([]T, bool, error) {
		start := in.Checkpoint()
		match := make([]T, 0)
		for {
			m, ok, err := parser(in)
			if err != nil {
				in.Restore(start)
				return nil, false, err
			}
			if !ok {
				in.Restore(start)
				return nil, false, nil
			}
			match = append(match, m)

			beforeDelimiter := in.Checkpoint()
			_, ok, err = delimiter(in)
			if err != nil {
				in.Restore(start)
				return nil, false, err
			}
			if ok {
				in.Restore(beforeDelimiter)
				return match, true, nil
			}
		}
	}
}

// UntilEOF matched until the delimiter is found or the end of the input is reached.
// This requires at least once parse to be successful before the end of the input.
// See tests for examples.
func UntilEOF[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return Until(parser, Or(delimiter, EOF[T]()))
}
