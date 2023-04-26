package set_test

import (
	"github.com/bang-go/util/set"
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	s := set.NewSet[int](1, 2, 3, 3)
	log.Println(s)
	s.Add(3, 4)
	log.Println(s)
	s.Remove(4)
	log.Println(s)
	log.Println(s.Len())
	log.Println(s.Slice())
	log.Println(s.Contains(2, 4))
	log.Println(s.Contains(3))
	log.Println(s.Clone())
	log.Println(s.Equal(s.Clone()))
	iterator := s.Iterator()
	for {
		if !iterator.HasNext() {
			break
		}
		log.Println(iterator.Next())
	}

	log.Println(s.Union(set.NewSet(5, 6, 7)))

	s.Clear()
	log.Println(s.Len())

}

func TestSetIterator(t *testing.T) {
	s := set.NewSetWithSize[string](3)
	s.Add("a", "b", "c")
	iterator := s.Iterator()
	for {
		if !iterator.HasNext() {
			break
		}
		log.Println(iterator.Next())
	}
}
