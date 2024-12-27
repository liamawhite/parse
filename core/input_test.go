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

	. "github.com/liamawhite/parse/core"
	"github.com/stretchr/testify/assert"
)

const input = `Some words
and some more words
and even more words`

func TestInput(t *testing.T) {
	i := NewInput(input)
	start := i.Checkpoint()

	t.Run("Peek forward", func(t *testing.T) {
		s, ok := i.Peek(10)
		assert.True(t, ok)
		assert.Equal(t, "Some words", s)
	})
	t.Run("Peek forward single", func(t *testing.T) {
		s, ok := i.Peek(1)
		assert.True(t, ok)
		assert.Equal(t, "S", s)
	})
	t.Run("Peek zero", func(t *testing.T) {
		s, ok := i.Peek(0)
		assert.True(t, ok)
		assert.Equal(t, "", s)
	})
	t.Run("Peek backward single", func(t *testing.T) {
		i.Take(1)
		s, ok := i.Peek(-1)
		assert.True(t, ok)
		assert.Equal(t, "S", s)
		i.Restore(start)
	})
	t.Run("Peek backward", func(t *testing.T) {
		i.Take(10) // need to move the index so we can peek backwards
		s, ok := i.Peek(-10)
		assert.True(t, ok)
		assert.Equal(t, "Some words", s)
		i.Restore(start)
	})
	t.Run("Peek backward beyond start", func(t *testing.T) {
		s, ok := i.Peek(-5)
		assert.False(t, ok)
		assert.Equal(t, "", s)
	})
	t.Run("Peek forward beyond end", func(t *testing.T) {
		s, ok := i.Peek(100)
		assert.False(t, ok)
		assert.Equal(t, "", s)
	})
	t.Run("Take forward", func(t *testing.T) {
		s, ok := i.Take(10)
		assert.True(t, ok)
		assert.Equal(t, "Some words", s)
		i.Restore(start)
	})
	t.Run("Take forward single", func(t *testing.T) {
		s, ok := i.Take(1)
		assert.True(t, ok)
		assert.Equal(t, "S", s)
		i.Restore(start)
	})
	t.Run("Take forward zero", func(t *testing.T) {
		s, ok := i.Take(0)
		assert.True(t, ok)
		assert.Equal(t, "", s)
	})
	t.Run("Take forward beyond end", func(t *testing.T) {
		s, ok := i.Take(100)
		assert.False(t, ok)
		assert.Equal(t, "", s)
	})
	t.Run("Take backward fails", func(t *testing.T) {
		s, ok := i.Take(-1)
		assert.False(t, ok)
		assert.Equal(t, "", s)
	})

}
