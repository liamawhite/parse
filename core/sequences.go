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

// SequenceOf2 parses two values in order or rolls back the input.
func SequenceOf2[A any, B any](a Parser[A], b Parser[B]) Parser[Tuple2[A, B]] {
	return func(in Input) (Tuple2[A, B], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple2[A, B]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple2[A, B]{}, false, errB
		}

		return tuple2[A, B]{A: matchA, B: matchB}, true, nil
	}
}

// SequenceOf3 parses three values in order or rolls back the input.
func SequenceOf3[A any, B any, C any](a Parser[A], b Parser[B], c Parser[C]) Parser[Tuple3[A, B, C]] {
	return func(in Input) (Tuple3[A, B, C], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple3[A, B, C]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple3[A, B, C]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple3[A, B, C]{}, false, errC
		}

		return tuple3[A, B, C]{A: matchA, B: matchB, C: matchC}, true, nil
	}
}

// SequenceOf4 parses four values in order or rolls back the input.
func SequenceOf4[A any, B any, C any, D any](a Parser[A], b Parser[B], c Parser[C], d Parser[D]) Parser[Tuple4[A, B, C, D]] {
	return func(in Input) (Tuple4[A, B, C, D], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple4[A, B, C, D]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple4[A, B, C, D]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple4[A, B, C, D]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple4[A, B, C, D]{}, false, errD
		}

		return tuple4[A, B, C, D]{A: matchA, B: matchB, C: matchC, D: matchD}, true, nil
	}
}

// SequenceOf5 parses five values in order or rolls back the input.
func SequenceOf5[A any, B any, C any, D any, E any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E]) Parser[Tuple5[A, B, C, D, E]] {
	return func(in Input) (Tuple5[A, B, C, D, E], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple5[A, B, C, D, E]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple5[A, B, C, D, E]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple5[A, B, C, D, E]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple5[A, B, C, D, E]{}, false, errD
		}

		matchE, okE, errE := e(in)
		if errE != nil || !okE {
			in.Restore(start)
			return tuple5[A, B, C, D, E]{}, false, errE
		}

		return tuple5[A, B, C, D, E]{A: matchA, B: matchB, C: matchC, D: matchD, E: matchE}, true, nil
	}
}

// SequenceOf6 parses six values in order or rolls back the input.
func SequenceOf6[A any, B any, C any, D any, E any, F any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F]) Parser[Tuple6[A, B, C, D, E, F]] {
	return func(in Input) (Tuple6[A, B, C, D, E, F], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errD
		}

		matchE, okE, errE := e(in)
		if errE != nil || !okE {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errE
		}

		matchF, okF, errF := f(in)
		if errF != nil || !okF {
			in.Restore(start)
			return tuple6[A, B, C, D, E, F]{}, false, errF
		}

		return tuple6[A, B, C, D, E, F]{A: matchA, B: matchB, C: matchC, D: matchD, E: matchE, F: matchF}, true, nil
	}
}

// SequenceOf7 parses seven values in order or rolls back the input.
func SequenceOf7[A any, B any, C any, D any, E any, F any, G any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G]) Parser[Tuple7[A, B, C, D, E, F, G]] {
	return func(in Input) (Tuple7[A, B, C, D, E, F, G], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errD
		}

		matchE, okE, errE := e(in)
		if errE != nil || !okE {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errE
		}

		matchF, okF, errF := f(in)
		if errF != nil || !okF {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errF
		}

		matchG, okG, errG := g(in)
		if errG != nil || !okG {
			in.Restore(start)
			return tuple7[A, B, C, D, E, F, G]{}, false, errG
		}

		return tuple7[A, B, C, D, E, F, G]{A: matchA, B: matchB, C: matchC, D: matchD, E: matchE, F: matchF, G: matchG}, true, nil
	}
}

// SequenceOf8 parses eight values in order or rolls back the input.
func SequenceOf8[A any, B any, C any, D any, E any, F any, G any, H any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G], h Parser[H]) Parser[Tuple8[A, B, C, D, E, F, G, H]] {
	return func(in Input) (Tuple8[A, B, C, D, E, F, G, H], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errD
		}

		matchE, okE, errE := e(in)
		if errE != nil || !okE {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errE
		}

		matchF, okF, errF := f(in)
		if errF != nil || !okF {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errF
		}

		matchG, okG, errG := g(in)
		if errG != nil || !okG {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errG
		}

		matchH, okH, errH := h(in)
		if errH != nil || !okH {
			in.Restore(start)
			return tuple8[A, B, C, D, E, F, G, H]{}, false, errH
		}

		return tuple8[A, B, C, D, E, F, G, H]{A: matchA, B: matchB, C: matchC, D: matchD, E: matchE, F: matchF, G: matchG, H: matchH}, true, nil
	}
}

// SequenceOf9 parses nine values in order or rolls back the input.
func SequenceOf9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a Parser[A], b Parser[B], c Parser[C], d Parser[D], e Parser[E], f Parser[F], g Parser[G], h Parser[H], i Parser[I]) Parser[Tuple9[A, B, C, D, E, F, G, H, I]] {
	return func(in Input) (Tuple9[A, B, C, D, E, F, G, H, I], bool, error) {
		start := in.Checkpoint()
		matchA, okA, errA := a(in)
		if errA != nil || !okA {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errA
		}

		matchB, okB, errB := b(in)
		if errB != nil || !okB {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errB
		}

		matchC, okC, errC := c(in)
		if errC != nil || !okC {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errC
		}

		matchD, okD, errD := d(in)
		if errD != nil || !okD {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errD
		}

		matchE, okE, errE := e(in)
		if errE != nil || !okE {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errE
		}

		matchF, okF, errF := f(in)
		if errF != nil || !okF {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errF
		}

		matchG, okG, errG := g(in)
		if errG != nil || !okG {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errG
		}

		matchH, okH, errH := h(in)
		if errH != nil || !okH {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errH
		}

		matchI, okI, errI := i(in)
		if errI != nil || !okI {
			in.Restore(start)
			return tuple9[A, B, C, D, E, F, G, H, I]{}, false, errI
		}

		return tuple9[A, B, C, D, E, F, G, H, I]{A: matchA, B: matchB, C: matchC, D: matchD, E: matchE, F: matchF, G: matchG, H: matchH, I: matchI}, true, nil
	}
}
