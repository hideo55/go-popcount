// +build amd64

package popcount

func havePOPCNT() bool
func popcnt64ASM(x uint64) uint64

var popcnt64 = popcnt64ASM

func init() {
	if !havePOPCNT() {
		popcnt64 = popcnt64Go
	}
}
