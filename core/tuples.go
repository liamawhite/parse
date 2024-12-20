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

type tuple2[A, B any] struct {
	A A
	B B
}

func NewTuple2[A, B any](a A, b B) tuple2[A, B] {
	return tuple2[A, B]{A: a, B: b}
}

type Tuple2[A, B any] interface {
	Values() (A, B)
}

func (t tuple2[A, B]) Values() (A, B) {
	return t.A, t.B
}

type tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

func (t tuple3[A, B, C]) Values() (A, B, C) {
	return t.A, t.B, t.C
}

func NewTuple3[A, B, C any](a A, b B, c C) tuple3[A, B, C] {
	return tuple3[A, B, C]{A: a, B: b, C: c}
}

type Tuple3[A, B, C any] interface {
	Values() (A, B, C)
}

type tuple4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

func NewTuple4[A, B, C, D any](a A, b B, c C, d D) tuple4[A, B, C, D] {
	return tuple4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

func (t tuple4[A, B, C, D]) Values() (A, B, C, D) {
	return t.A, t.B, t.C, t.D
}

type Tuple4[A, B, C, D any] interface {
	Values() (A, B, C, D)
}

type tuple5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func NewTuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) tuple5[A, B, C, D, E] {
	return tuple5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

func (t tuple5[A, B, C, D, E]) Values() (A, B, C, D, E) {
	return t.A, t.B, t.C, t.D, t.E
}

type Tuple5[A, B, C, D, E any] interface {
	Values() (A, B, C, D, E)
}

type tuple6[A, B, C, D, E, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func NewTuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) tuple6[A, B, C, D, E, F] {
	return tuple6[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

func (t tuple6[A, B, C, D, E, F]) Values() (A, B, C, D, E, F) {
	return t.A, t.B, t.C, t.D, t.E, t.F
}

type Tuple6[A, B, C, D, E, F any] interface {
	Values() (A, B, C, D, E, F)
}

type tuple7[A, B, C, D, E, F, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

func NewTuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) tuple7[A, B, C, D, E, F, G] {
	return tuple7[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

func (t tuple7[A, B, C, D, E, F, G]) Values() (A, B, C, D, E, F, G) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G
}

type Tuple7[A, B, C, D, E, F, G any] interface {
	Values() (A, B, C, D, E, F, G)
}

type tuple8[A, B, C, D, E, F, G, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

func NewTuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) tuple8[A, B, C, D, E, F, G, H] {
	return tuple8[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

func (t tuple8[A, B, C, D, E, F, G, H]) Values() (A, B, C, D, E, F, G, H) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H
}

type Tuple8[A, B, C, D, E, F, G, H any] interface {
	Values() (A, B, C, D, E, F, G, H)
}

type tuple9[A, B, C, D, E, F, G, H, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

func NewTuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) tuple9[A, B, C, D, E, F, G, H, I] {
	return tuple9[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

func (t tuple9[A, B, C, D, E, F, G, H, I]) Values() (A, B, C, D, E, F, G, H, I) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I
}

type Tuple9[A, B, C, D, E, F, G, H, I any] interface {
	Values() (A, B, C, D, E, F, G, H, I)
}
