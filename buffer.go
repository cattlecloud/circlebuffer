// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package circlebuffer

import "iter"

type Array[T any] interface {
	All() iter.Seq[T]
	Insert(T)
	Size() int
	Empty() bool
}

func New[T any](capacity int) Array[T] {
	return &array[T]{
		data: make([]T, capacity),
	}
}

type array[T any] struct {
	data []T
	head int
	size int
}

func (b *array[T]) Size() int {
	return b.size
}

func (b *array[T]) Empty() bool {
	return b.Size() == 0
}

func (b *array[T]) Insert(item T) {
	b.data[b.head] = item
	if b.size < len(b.data) {
		b.size++
	}
	b.head = (b.head + 1) % len(b.data)
}

func (b *array[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if b.size == 0 {
			return
		}

		start := (b.head - b.size + len(b.data)) % len(b.data)
		for i := 0; i < b.size; i++ {
			index := (start + i) % len(b.data)
			if !yield(b.data[index]) {
				return
			}
		}
	}
}
