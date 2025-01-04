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

import (
	"strings"
)

type Input interface {
	Peek(n int) (string, bool)
	Take(n int) (string, bool)
	Checkpoint() int
	Restore(checkpoint int)
	Debug() string
}

type input struct {
	s     string
	index int
}

func NewInput(s string) Input {
	return &input{
		s: s,
	}
}

func (i *input) Peek(n int) (s string, ok bool) {
	if i.index+n > len(i.s) || i.index+n < 0 {
		return
	}
	if n < 0 {
		return i.s[i.index+n : i.index], true
	}
	return i.s[i.index : i.index+n], true
}

func (i *input) Take(n int) (s string, ok bool) {
	if i.index+n > len(i.s) {
		return
	}
	if n < 0 {
		return
	}
	from := i.index
	i.index += n
	return i.s[from:i.index], true
}

// Take a snapshot of the current parsing position
func (i *input) Checkpoint() int {
	return i.index
}

// Restore the parsing position to a previous snapshot
func (i *input) Restore(checkpoint int) {
	if checkpoint < 0 {
		checkpoint = 0
	}
	if checkpoint > len(i.s) {
		checkpoint = len(i.s)
	}
	i.index = checkpoint
}

// Outputs the full input string with a marker at the current parsing position
func (i *input) Debug() string {
	var s strings.Builder
	s.WriteString(i.s[:i.index])
	s.WriteString("☝️")
	s.WriteString(i.s[i.index:])
	return s.String()
}
