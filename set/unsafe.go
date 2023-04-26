package set

import (
	"fmt"
	"github.com/bang-go/util"
	"strings"
)

type unsafeMap[T comparable] map[T]struct{}

func newUnsafeMap[T comparable](size int) unsafeMap[T] {
	return make(map[T]struct{}, size)
}

func (u unsafeMap[T]) add(val ...T) {
	for _, elem := range val {
		u[elem] = struct{}{}
	}
}

func (u unsafeMap[T]) remove(val ...T) {
	for _, elem := range val {
		delete(u, elem)
	}
}

func (u unsafeMap[T]) contains(val ...T) bool {
	for _, elem := range val {
		if _, ok := u[elem]; !ok {
			return false
		}
	}
	return true
}

func (u unsafeMap[T]) clear() {
	//s.m = map[T]struct{}{}
	for item := range u {
		delete(u, item)
	}
}

func (u unsafeMap[T]) len() int {
	return len(u)
}

func (u unsafeMap[T]) slice() []T {
	list := make([]T, 0, u.len())
	for item := range u {
		list = append(list, item)
	}
	return list
}

func (u unsafeMap[T]) string() string {
	itemStr := make([]string, 0, u.len())
	for item := range u {
		itemStr = append(itemStr, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(itemStr, ","))
}

func (u unsafeMap[T]) clone() unsafeMap[T] {
	newMap := newUnsafeMap[T](u.len())
	newMap.add(u.slice()...)
	return newMap
}

func (u unsafeMap[T]) equal(other unsafeMap[T]) bool {
	if other.len() != u.len() {
		return false
	}
	for item := range other {
		if !other.contains(item) {
			return false
		}
	}
	return true
}

func (u unsafeMap[T]) union(other unsafeMap[T]) unsafeMap[T] {
	sLen := other.len()
	oLen := other.len()
	maxLen := util.If(sLen > oLen, sLen, oLen)
	newMap := newUnsafeMap[T](maxLen)

	for item := range u {
		newMap.add(item)
	}
	for item := range other {
		newMap.add(item)
	}
	return newMap
}
