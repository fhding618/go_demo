// go test -benchmem -bench ^BenchmarkNew_01$
// go test -benchmem -bench ^BenchmarkNew_02$

package singleton

import "testing"

func BenchmarkNew_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New_01()
	}
}

func BenchmarkNew_02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New_02()
	}
}
