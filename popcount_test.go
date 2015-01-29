package popcount

import (
    "testing"
)

type (
    testCase64 struct {
        n uint64
        x uint64
    }
)

var (
    cases64 = []testCase64 {
        { 64, 0xFFFFFFFFFFFFFFFF },
        { 0, 0x0 },
        { 1, 0x1 },
        { 1, 0x8000000000000000 },
        { 2, 0x8000000000000001 },
        { 4, 0xF0 },
        { 7, 0xFE },
        { 32, 0x123456789ABCDEF },
    }
)

func TestPopCount64(t *testing.T) {
    for _, tc := range cases64 {
        var count uint64 = Count(tc.x)
        if count != tc.n {
            t.Error("Expected", tc.n, "got", count)
        }
    }
}

type (
    testCaseSlice struct {
        n uint64
        s []uint64
    }
)

var (
    caseSlice = []testCaseSlice {
        { 120, []uint64{0xFF, 0xFFFF, 0xFFFFFFFF, 0xFFFFFFFFFFFFFFFF }, },
        { 17, []uint64{1,2,3,4,5,6,7,8,9,10}, },
    }
)

func TestPopCountSlice(t *testing.T) {
    for _, tc := range caseSlice {
        var count = CountSlice(tc.s)
        if count != tc.n {
            t.Error("Expected", tc.n, "got", count)
        }
    }
}
