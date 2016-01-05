package popcount

import (
	"testing"
	"testing/quick"
)

func TestPopCountCompare(t *testing.T) {
	if !asm {
		t.Skip()
	}

	for _, tc := range cases64 {
		if popcnt64ASM(tc.x) != popcnt64Go(tc.x) {
			t.Errorf("popcnt64ASM(%v) = %v; popcnt64Go(%v) = %d",
				tc.x, popcnt64ASM(tc.x), tc.x, popcnt64Go(tc.x))
		}
	}

	if err := quick.CheckEqual(popcnt64Go, popcnt64ASM, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkASM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt64ASM(0x123456789ABCDEF)
	}
}

func BenchmarkASMIndirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt64(0x123456789ABCDEF)
	}
}
