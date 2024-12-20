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

// Use unexported type alias so we control the interface and prevent people doing arbirary things e.g. use negative ints.
type checkpoint int

type Input interface {
	Peek(n int) (string, bool)
	Take(n int) (string, bool)
	Checkpoint() checkpoint
	Restore(cp checkpoint)
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
	if i.index+n > len(i.s) {
		return
	}
	if n < 0 {
		return i.s[i.index:], true
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
func (i *input) Checkpoint() checkpoint {
	return checkpoint(i.index)
}

// Restore the parsing position to a previous snapshot
func (i *input) Restore(cp checkpoint) {
	i.index = int(cp)
}
