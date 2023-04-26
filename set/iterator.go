package set

import "sync"

type Iterator[T comparable] interface {
	HasNext() bool
	Next() T
}

type iteratorWrapper[T comparable] struct {
	sync.RWMutex
	index int
	list  []T
}

func NewIterator[T comparable](list []T) Iterator[T] {
	return &iteratorWrapper[T]{list: list, index: 0}
}

func (i *iteratorWrapper[T]) HasNext() bool {
	return len(i.list) > i.index
}

func (i *iteratorWrapper[T]) Next() (t T) {
	i.RLock()
	defer i.RUnlock()
	if i.HasNext() {
		t = i.list[i.index]
		i.index++
		return
	}
	return
}
