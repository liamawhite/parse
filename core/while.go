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

// WhileNot matches while the given delimiter is not found.
// This differs from until as the delimiter is searched for BEFORE the first match.
func WhileNot[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return func(in Input) ([]T, bool, error) {
		start := in.Checkpoint()
		match := make([]T, 0)
		for {
			beforeDelimiter := in.Checkpoint()
			_, ok, err := delimiter(in)
			if err != nil {
				in.Restore(start)
				return nil, false, err
			}
			if ok {
				in.Restore(beforeDelimiter)
				return match, true, nil
			}

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
		}
	}
}

// WhileNotEOFOr matches while the given delimiter is not found or the end of the input is reached.
// This requires no successful parses before the delimiter or EOF is found.
// See tests for examples.
func WhileNotEOFOr[T, D any](parser Parser[T], delimiter Parser[D]) Parser[[]T] {
	return WhileNot(parser, Or(delimiter, EOF[T]()))
}
