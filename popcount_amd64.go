// +build amd64

package popcount

func havePOPCNT() bool
func popcnt64ASM(x uint64) uint64

var asm = havePOPCNT()

func popcnt64(x uint64) uint64 {
	if asm {
		return popcnt64ASM(x)
	}
	return popcnt64Go(x)
}
