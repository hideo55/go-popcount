// Package popcount is PopCount implementation for Go. 
// Using hardware POPCNT instruction if available it.
package popcount

// CountSlice counts number of non-zero bits in slice of 64bit unsigned integer.
func CountSlice(s []uint64) uint64 {
    count := uint64(0)
    for _, x := range s {
        count += popcnt64(x)
    }
    return count
}

// Count counts number of non-zero bits in 64bit unsigned integer.
func Count(x uint64) uint64 {
    return popcnt64(x)
}

func popcnt64Go(x uint64) uint64 {
    x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
    x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2);
    x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4);
    x *= 0x0101010101010101
    return ((x >> 56) & 0xFF)
}

