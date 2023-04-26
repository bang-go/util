package set

import (
	"sync"
)

type Set[T comparable] interface {
	Add(val ...T)
	Remove(val ...T)
	Clear()
	Contains(val ...T) bool
	String() string
	Slice() []T
	Len() int
	Clone() Set[T]
	Iterator() Iterator[T]
	Equal(other Set[T]) bool
	Union(other Set[T]) Set[T]
}

type setWrapper[T comparable] struct {
	sync.RWMutex
	m unsafeMap[T]
}

func NewSet[T comparable](val ...T) Set[T] {
	s := NewSetWithSize[T](len(val))
	s.Add(val...)
	return s
}

func NewSetWithSize[T comparable](size int) Set[T] {
	m := newUnsafeMap[T](size)
	return &setWrapper[T]{
		m: m,
	}
}

func (s *setWrapper[T]) Add(val ...T) {
	s.Lock()
	defer s.Unlock()
	s.m.add(val...)
}

func (s *setWrapper[T]) Remove(val ...T) {
	s.Lock()
	defer s.Unlock()
	s.m.remove(val...)
}

func (s *setWrapper[T]) Contains(val ...T) bool {
	s.RLock()
	defer s.RUnlock()
	return s.m.contains(val...)
}

func (s *setWrapper[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m.clear()
}

func (s *setWrapper[T]) Len() int {
	return s.m.len()
}

func (s *setWrapper[T]) Slice() []T {
	s.RLock()
	defer s.RUnlock()
	return s.m.slice()
}

func (s *setWrapper[T]) String() string {
	return s.m.string()
}

func (s *setWrapper[T]) Clone() Set[T] {
	s.Lock()
	defer s.Unlock()
	newSet := &setWrapper[T]{m: s.m.clone()}
	return newSet
}

func (s *setWrapper[T]) Iterator() Iterator[T] {
	return NewIterator[T](s.Slice())
}

func (s *setWrapper[T]) Equal(other Set[T]) bool {
	o := other.(*setWrapper[T])
	s.Lock()
	o.Lock()
	defer s.Unlock()
	defer o.Unlock()
	return o.m.equal(o.m)
}

func (s *setWrapper[T]) Union(other Set[T]) Set[T] {
	o := other.(*setWrapper[T])
	s.Lock()
	o.Lock()
	defer s.Unlock()
	defer o.Unlock()
	newSet := &setWrapper[T]{m: s.m.union(o.m)}
	return newSet
}
