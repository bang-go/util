package util_test

import (
	"github.com/bang-go/util"
	"log"
	"testing"
)

func TestStringRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		str := util.StringRand(10)
		log.Println(str)
	}
}
