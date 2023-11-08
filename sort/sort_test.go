package sort

import (
	"math/rand"
	"testing"
	"time"

	"algs4/typ"
	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 1},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			SelectionSort(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			InsertionSort(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			BubbleSort(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestShellSort(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			ShellSort(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			MergeSortInt(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestQuick(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sli := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			if tt.count > 1 {
				require.False(t, sli.IsSorted())
			}
			QuickSortInt(sli)
			require.True(t, sli.IsSorted())
		})
	}
}

func TestMaxPQ_IsMaxHeap(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := typ.IntSlice(GenRandomNums(40000000, 20000000))
			keys := make([]typ.Comparable, len(input))
			for i := range input {
				keys[i] = typ.ComparableInt(input[i])
			}
			maxPQ := NewMaxPQWithKeys(keys)
			require.True(t, maxPQ.IsMaxHeap())
		})
	}
}

func TestMaxPQ(t *testing.T) {
	tests := []struct {
		name   string
		keys   []int
		max    typ.Comparable
		top5   []int
		top    typ.Comparable
		nStart int
		nFinal int
	}{
		{
			name:   "case1",
			keys:   []int{23, 2, 32, 43, 54, 654, 432, -1, 34534, 239},
			max:    typ.ComparableInt(34534),
			top5:   []int{34534, 654, 432, 239, 54},
			top:    typ.ComparableInt(43),
			nStart: 10,
			nFinal: 4,
		},
		{
			name:   "case2",
			keys:   []int{1, 2, 3, 4, 5, 6},
			max:    typ.ComparableInt(6),
			top5:   []int{6, 5, 4, 3, 2},
			top:    typ.ComparableInt(1),
			nStart: 6,
			nFinal: 0,
		},
		{
			name:   "case3",
			keys:   []int{99, 9, -9, 342, 111, 111, -10, -8, 0, 111, 111, 99, 0, 0, 0, -8, -8},
			max:    typ.ComparableInt(342),
			top5:   []int{342, 111, 111, 111, 111},
			top:    typ.ComparableInt(99),
			nStart: 17,
			nFinal: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := make([]typ.Comparable, len(tt.keys))
			for i := range tt.keys {
				keys[i] = typ.ComparableInt(tt.keys[i])
			}
			pq := NewMaxPQWithKeys(keys)

			require.Equal(t, tt.nStart, pq.Size())

			require.True(t, pq.IsMaxHeap())

			require.Equal(t, tt.max, pq.Max())

			top5, err := pq.TopK(5)
			top5Int := make([]int, len(top5))
			for i := range top5 {
				top5Int[i] = int(top5[i].(typ.ComparableInt))
			}
			require.NoError(t, err)
			require.Equal(t, tt.top5, top5Int)

			top := pq.DelMax()
			require.Equal(t, tt.top, top)

			require.Equal(t, tt.nFinal, pq.Size())
		})
	}
}

func TestMinPQ_IsMinHeap(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := typ.IntSlice(GenRandomNums(40000000, 20000000))
			keys := make([]typ.Comparable, len(input))
			for i := range input {
				keys[i] = typ.ComparableInt(input[i])
			}
			maxPQ := NewMinPQWithKeys(keys)
			require.True(t, maxPQ.IsMinHeap())
		})
	}
}

func TestMinPQ(t *testing.T) {
	tests := []struct {
		name   string
		keys   []int
		min    typ.Comparable
		top5   []int
		top    typ.Comparable
		nStart int
		nFinal int
	}{
		{
			name:   "case1",
			keys:   []int{23, 2, 32, 43, 54, 654, 432, -1, 34534, 239},
			min:    typ.ComparableInt(-1),
			top5:   []int{-1, 2, 23, 32, 43},
			top:    typ.ComparableInt(54),
			nStart: 10,
			nFinal: 4,
		},
		{
			name:   "case2",
			keys:   []int{1, 2, 3, 4, 5, 6},
			min:    typ.ComparableInt(1),
			top5:   []int{1, 2, 3, 4, 5},
			top:    typ.ComparableInt(6),
			nStart: 6,
			nFinal: 0,
		},
		{
			name:   "case3",
			keys:   []int{99, 9, -9, 342, 111, 111, -10, -8, 0, 111, 111, 99, 0, 0, 0, -8, -8},
			min:    typ.ComparableInt(-10),
			top5:   []int{-10, -9, -8, -8, -8},
			top:    typ.ComparableInt(0),
			nStart: 17,
			nFinal: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := make([]typ.Comparable, len(tt.keys))
			for i := range tt.keys {
				keys[i] = typ.ComparableInt(tt.keys[i])
			}
			pq := NewMinPQWithKeys(keys)

			require.Equal(t, tt.nStart, pq.Size())

			require.True(t, pq.IsMinHeap())

			require.Equal(t, tt.min, pq.Min())

			top5, err := pq.MinK(5)
			top5Int := make([]int, len(top5))
			for i := range top5 {
				top5Int[i] = int(top5[i].(typ.ComparableInt))
			}
			require.NoError(t, err)
			require.Equal(t, tt.top5, top5Int)

			top := pq.DelMin()
			require.Equal(t, tt.top, top)

			require.Equal(t, tt.nFinal, pq.Size())
		})
	}
}

func TestIndexMinPQ(t *testing.T) {
	// TODO: Add test cases.
}

func TestHeap(t *testing.T) {
	var tests = []struct {
		name       string
		max, count int
	}{
		{"case1", 65536, 0},
		{"case2", 65536, 10},
		{"case3", 65536, 100},
		{"case4", 65536, 1000},
		{"case5", 65536, 10000},
		{"case6", 65536, 100000},
		{"case7", 1024, 1000},
		{"case8", 1024, 10000},
		{"case9", 1024, 100000},
		{"case10", 6553600, 1000},
		{"case11", 6553600, 10000},
		{"case12", 6553600, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := typ.IntSlice(GenRandomNums(tt.max, tt.count))
			keys := make([]typ.Comparable, len(input))
			for i := range input {
				keys[i] = typ.ComparableInt(input[i])
			}
			h := Heap{pq: keys}
			h.Sort()
			require.True(t, h.IsSorted())
		})
	}
}

func GenRandomNums(max, count int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := make([]int, count)
	for i := 0; i < count; i++ {
		ret[i] = rand.Intn(max)
	}
	return ret
}
