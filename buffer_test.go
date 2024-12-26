// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package circlebuffer

import (
	"testing"

	"github.com/shoenig/test/must"
)

func testSequence[T any](t *testing.T, a Array[T], exp []T) {
	t.Helper()

	result := make([]T, 0, a.Size())
	for item := range a.All() {
		result = append(result, item)
	}
	must.Eq(t, exp, result)
}

func TestArray(t *testing.T) {
	t.Parallel()

	buf := New[string](3)
	must.Empty(t, buf)
	testSequence(t, buf, []string{})

	buf.Insert("one")
	must.Size(t, 1, buf)
	testSequence(t, buf, []string{"one"})

	buf.Insert("two")
	must.Size(t, 2, buf)
	testSequence(t, buf, []string{"one", "two"})

	buf.Insert("three")
	must.Size(t, 3, buf)
	testSequence(t, buf, []string{"one", "two", "three"})

	buf.Insert("four")
	must.Size(t, 3, buf)
	testSequence(t, buf, []string{"two", "three", "four"})

	buf.Insert("five")
	must.Size(t, 3, buf)
	testSequence(t, buf, []string{"three", "four", "five"})
}
