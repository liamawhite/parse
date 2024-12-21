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

// Or returns successful if either of the parsers are successful.
// It returns as soon as one of the parsers are successful or rolls back when none are.
func Or[A any, B any](a Parser[A], b Parser[B]) Parser[Tuple2[Match[A], Match[B]]] {
	return func(in Input) (Tuple2[Match[A], Match[B]], bool, error) {
		var res tuple2[Match[A], Match[B]]

		matchA, okA, errA := a(in)
		if errA != nil {
			return res, false, errA
		}
		res.A = NewMatch(matchA, okA)
		if okA {
			return res, true, nil
		}

		matchB, okB, errB := b(in)
		if errB != nil {
			return res, false, errB
		}
		res.B = NewMatch(matchB, okB)
		if okB {
			return res, true, nil
		}

		return res, false, nil
	}
}
