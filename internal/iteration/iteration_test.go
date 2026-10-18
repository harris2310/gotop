package iteration

import "testing"

func BenchmarkIteration(b *testing.B) {
	for b.Loop() {
		RunIteration()
	}
}
