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

// Any looks for matches in the given parsers, returning the first match or rolls back the input if no match is found.
func Any[T any](parsers ...Parser[T]) Parser[T] {
	return func(in Input) (T, bool, error) {
		for _, parser := range parsers {
			match, ok, err := parser(in)
			if err != nil || ok {
				return match, true, err
			}
		}
		var t T
		return t, false, nil
	}
}
