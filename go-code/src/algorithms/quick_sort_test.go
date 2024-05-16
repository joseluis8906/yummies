package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/go-code/src/algorithms"
	"github.com/joseluis8906/go-code/src/pkg/cmp"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()

	want := []int{0, 0, 0, 1, 2, 2, 3, 4, 5, 6, 7, 0, 0, 0}
	got := []int{0, 0, 0, 2, 4, 5, 7, 1, 2, 3, 6, 0, 0, 0}

	algorithms.QuickSort(got, 3, 10)

	if !cmp.Equal(got, want) {
		t.Errorf("QuickSort() = %v, want %v\n%v", got, want, cmp.Diff(want, got))
	}
}

func BenchmarkQuickSort(b *testing.B) {
	testCases := []struct {
		name  string
		input []int
		p     int
		r     int
	}{
		{
			name:  "Small",
			input: []int{9, 6, 2},
			p:     0,
			r:     2,
		},
		{
			name:  "Middle",
			input: []int{10, 70, 20, 5, 33, 22},
			p:     0,
			r:     5,
		},
		{
			name:  "Big",
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			p:     0,
			r:     8,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algorithms.QuickSort(tc.input, tc.p, tc.r)
			}
		})
	}
}
