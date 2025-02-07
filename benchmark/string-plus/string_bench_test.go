package stringplus

import (
	"fmt"
	"strings"
	"testing"
)

func ff(p ...interface{}) {

}

func BenchmarkSimpleTest(b *testing.B) {
	str10 := genStr(10)
	str100 := genStr(100)
	for i := 0; i < b.N; i++ {
		var buf []byte
		ff(buf)
		buf = append(buf, str10...)
		buf = append(buf, str100...)
	}
}

func BenchmarkStrPlus(b *testing.B) {
	for _, tc := range testcase {
		b.Run(tc.name, func(b *testing.B) {
			p := genStr(tc.len)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var s string
				for j := 0; j < int(tc.times); j++ {
					s += p
				}
			}
		})
	}
}

func BenchmarkSprintf(b *testing.B) {
	for _, tc := range testcase {
		b.Run(tc.name, func(b *testing.B) {
			p := genStr(tc.len)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var s string
				for j := 0; j < int(tc.times); j++ {
					s = fmt.Sprintf("%s%s", s, p)
				}
			}
		})
	}
}

func BenchmarkBuilder(b *testing.B) {
	for _, tc := range testcase {
		b.Run(tc.name, func(b *testing.B) {
			p := genStr(tc.len)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bd := strings.Builder{}
				bd.Grow(int(tc.times) * int(tc.len))
				for j := 0; j < int(tc.times); j++ {
					bd.WriteString(p)
				}
				_ = bd.String()
			}
		})
	}
}
