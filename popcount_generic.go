// +build !amd64

package popcount

func popcnt64(x uint64) uint64 {
    return popcnt64Go(x)
}
