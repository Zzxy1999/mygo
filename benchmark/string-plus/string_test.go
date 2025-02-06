package stringplus

import (
	"fmt"
	"strings"
	"testing"
)

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
				for j := 0; j < int(tc.times); j++ {
					bd.WriteString(p)
				}
				_ = bd.String()
			}
		})
	}
}
