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

package core_test

import (
	"testing"

	"github.com/liamawhite/parse/core"
	. "github.com/liamawhite/parse/test"
)

func TestSequence2(t *testing.T) {
	tests := []ParserTest[core.Tuple2[string, string]]{
		{
			Name:           "no match",
			Input:          "AB",
			Parser:         core.SequenceOf2(core.String("1"), core.String("B")),
			ExpectedMatch:  core.NewTuple2("", ""),
			ExpectedOK:     false,
			RemainingInput: "AB",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "AB",
			Parser:         core.SequenceOf2(NaughtyParser[string](), core.String("B")),
			ExpectedMatch:  core.NewTuple2("", ""),
			ExpectedOK:     false,
			RemainingInput: "AB",
		},
		{
			Name:           "partial match",
			Input:          "AB",
			Parser:         core.SequenceOf2(core.String("A"), core.String("2")),
			ExpectedMatch:  core.NewTuple2("", ""),
			ExpectedOK:     false,
			RemainingInput: "AB",
		},
		{
			Name:          "match",
			Input:         "AB",
			Parser:        core.SequenceOf2(core.String("A"), core.String("B")),
			ExpectedMatch: core.NewTuple2("A", "B"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence3(t *testing.T) {
	tests := []ParserTest[core.Tuple3[string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABC",
			Parser:         core.SequenceOf3(core.String("1"), core.String("B"), core.String("3")),
			ExpectedMatch:  core.NewTuple3("", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABC",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABC",
			Parser:         core.SequenceOf3(NaughtyParser[string](), core.String("B"), core.String("C")),
			ExpectedMatch:  core.NewTuple3("", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABC",
		},
		{
			Name:           "partial match",
			Input:          "ABC",
			Parser:         core.SequenceOf3(core.String("A"), core.String("2"), core.String("C")),
			ExpectedMatch:  core.NewTuple3("", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABC",
		},
		{
			Name:          "match",
			Input:         "ABC",
			Parser:        core.SequenceOf3(core.String("A"), core.String("B"), core.String("C")),
			ExpectedMatch: core.NewTuple3("A", "B", "C"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence4(t *testing.T) {
	tests := []ParserTest[core.Tuple4[string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCD",
			Parser:         core.SequenceOf4(core.String("1"), core.String("B"), core.String("3"), core.String("D")),
			ExpectedMatch:  core.NewTuple4("", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCD",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCD",
			Parser:         core.SequenceOf4(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D")),
			ExpectedMatch:  core.NewTuple4("", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCD",
		},
		{
			Name:           "partial match",
			Input:          "ABCD",
			Parser:         core.SequenceOf4(core.String("A"), core.String("2"), core.String("C"), core.String("D")),
			ExpectedMatch:  core.NewTuple4("", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCD",
		},
		{
			Name:          "match",
			Input:         "ABCD",
			Parser:        core.SequenceOf4(core.String("A"), core.String("B"), core.String("C"), core.String("D")),
			ExpectedMatch: core.NewTuple4("A", "B", "C", "D"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence5(t *testing.T) {
	tests := []ParserTest[core.Tuple5[string, string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCDE",
			Parser:         core.SequenceOf5(core.String("1"), core.String("B"), core.String("3"), core.String("D"), core.String("E")),
			ExpectedMatch:  core.NewTuple5("", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDE",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCDE",
			Parser:         core.SequenceOf5(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D"), core.String("E")),
			ExpectedMatch:  core.NewTuple5("", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDE",
		},
		{
			Name:           "partial match",
			Input:          "ABCDE",
			Parser:         core.SequenceOf5(core.String("A"), core.String("2"), core.String("C"), core.String("D"), core.String("E")),
			ExpectedMatch:  core.NewTuple5("", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDE",
		},
		{
			Name:          "match",
			Input:         "ABCDE",
			Parser:        core.SequenceOf5(core.String("A"), core.String("B"), core.String("C"), core.String("D"), core.String("E")),
			ExpectedMatch: core.NewTuple5("A", "B", "C", "D", "E"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence6(t *testing.T) {
	tests := []ParserTest[core.Tuple6[string, string, string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCDEF",
			Parser:         core.SequenceOf6(core.String("1"), core.String("B"), core.String("3"), core.String("D"), core.String("E"), core.String("F")),
			ExpectedMatch:  core.NewTuple6("", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCDEF",
			Parser:         core.SequenceOf6(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F")),
			ExpectedMatch:  core.NewTuple6("", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:           "partial match",
			Input:          "ABCDEF",
			Parser:         core.SequenceOf6(core.String("A"), core.String("2"), core.String("C"), core.String("D"), core.String("E"), core.String("F")),
			ExpectedMatch:  core.NewTuple6("", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEF",
		},
		{
			Name:          "match",
			Input:         "ABCDEF",
			Parser:        core.SequenceOf6(core.String("A"), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F")),
			ExpectedMatch: core.NewTuple6("A", "B", "C", "D", "E", "F"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence7(t *testing.T) {
	tests := []ParserTest[core.Tuple7[string, string, string, string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCDEFG",
			Parser:         core.SequenceOf7(core.String("1"), core.String("B"), core.String("3"), core.String("D"), core.String("E"), core.String("F"), core.String("G")),
			ExpectedMatch:  core.NewTuple7("", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFG",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCDEFG",
			Parser:         core.SequenceOf7(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G")),
			ExpectedMatch:  core.NewTuple7("", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFG",
		},
		{
			Name:           "partial match",
			Input:          "ABCDEFG",
			Parser:         core.SequenceOf7(core.String("A"), core.String("2"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G")),
			ExpectedMatch:  core.NewTuple7("", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFG",
		},
		{
			Name:          "match",
			Input:         "ABCDEFG",
			Parser:        core.SequenceOf7(core.String("A"), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G")),
			ExpectedMatch: core.NewTuple7("A", "B", "C", "D", "E", "F", "G"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence8(t *testing.T) {
	tests := []ParserTest[core.Tuple8[string, string, string, string, string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCDEFGH",
			Parser:         core.SequenceOf8(core.String("1"), core.String("B"), core.String("3"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H")),
			ExpectedMatch:  core.NewTuple8("", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGH",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCDEFGH",
			Parser:         core.SequenceOf8(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H")),
			ExpectedMatch:  core.NewTuple8("", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGH",
		},
		{
			Name:           "partial match",
			Input:          "ABCDEFGH",
			Parser:         core.SequenceOf8(core.String("A"), core.String("2"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H")),
			ExpectedMatch:  core.NewTuple8("", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGH",
		},
		{
			Name:          "match",
			Input:         "ABCDEFGH",
			Parser:        core.SequenceOf8(core.String("A"), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H")),
			ExpectedMatch: core.NewTuple8("A", "B", "C", "D", "E", "F", "G", "H"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}

func TestSequence9(t *testing.T) {
	tests := []ParserTest[core.Tuple9[string, string, string, string, string, string, string, string, string]]{
		{
			Name:           "no match",
			Input:          "ABCDEFGHI",
			Parser:         core.SequenceOf9(core.String("1"), core.String("B"), core.String("3"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H"), core.String("I")),
			ExpectedMatch:  core.NewTuple9("", "", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGHI",
		},
		{
			Name:           "no match rolls back input even if one of the parsers consumed input",
			Input:          "ABCDEFGHI",
			Parser:         core.SequenceOf9(NaughtyParser[string](), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H"), core.String("I")),
			ExpectedMatch:  core.NewTuple9("", "", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGHI",
		},
		{
			Name:           "partial match",
			Input:          "ABCDEFGHI",
			Parser:         core.SequenceOf9(core.String("A"), core.String("2"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H"), core.String("I")),
			ExpectedMatch:  core.NewTuple9("", "", "", "", "", "", "", "", ""),
			ExpectedOK:     false,
			RemainingInput: "ABCDEFGHI",
		},
		{
			Name:          "match",
			Input:         "ABCDEFGHI",
			Parser:        core.SequenceOf9(core.String("A"), core.String("B"), core.String("C"), core.String("D"), core.String("E"), core.String("F"), core.String("G"), core.String("H"), core.String("I")),
			ExpectedMatch: core.NewTuple9("A", "B", "C", "D", "E", "F", "G", "H", "I"),
			ExpectedOK:    true,
		},
	}
	RunTests(t, tests)
}
