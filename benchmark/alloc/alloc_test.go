package alloc

import (
	"math/rand"
	"strconv"
	"testing"
)

func retStr() string {
	s := strconv.Itoa(rand.Int())
	return s
}

func BenchmarkStringCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = retStr()
	}
}
