// +build amd64

package popcount

func havePOPCNT() bool

var popcnt = havePOPCNT()

func popcnt64ASM(x uint64) uint64

func popcnt64(x uint64) uint64 {
	if popcnt {
		return popcnt64ASM(x)
	}
	return popcnt64Go(x)
}
