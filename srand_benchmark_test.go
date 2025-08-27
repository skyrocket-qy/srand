package srand

import "testing"

func BenchmarkStringFromCharset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StringFromCharset(16, CharsetAlphaNum)
		if err != nil {
			b.Fatal(err)
		}
	}
}
