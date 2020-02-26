package sort

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	test(t, Select)
}

func test(t *testing.T, s func(sort.Interface)) {
	tests := [][]int{
		{},
		{1, 2, 3, 4},
		{1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 4, 2, 5},
		{5, 3, 2, 4, 1},
		{400, 250, 249, 251, -777, 392, 42, 24, 70, 1, 24, 50},
		{-10, -27, -9, 0, 0, 0, 315, 314, 0, 91, 1, 0, -1},
	}
	for i, tt := range tests {
		t.Run("#"+strconv.Itoa(i), func(t *testing.T) {
			expect := make([]int, len(tt))
			copy(expect, tt)
			sort.Ints(expect)
			s(sort.IntSlice(tt))
			assert.EqualValues(t, expect, tt)
		})
	}
}

func BenchmarkSelect_Students(b *testing.B) {
	if data == nil {
		b.Skip()
	}
	for i := 0; i < b.N; i++ {
		Select(data)
	}
}
