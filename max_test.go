package window_test

import (
	"testing"

	"github.com/julz/window"
)

func TestWindowMax(t *testing.T) {
	tests := []struct {
		name      string
		values    []float64
		indexFunc func(int) int
		expect    float64
	}{{
		name:   "single value",
		values: []float64{1},
		expect: 1,
	}, {
		name:   "ascending values",
		values: []float64{1, 2},
		expect: 2,
	}, {
		name:   "descending values",
		values: []float64{2, 1},
		expect: 2,
	}, {
		name:   "up, down, up",
		values: []float64{1, 2, 1},
		expect: 2,
	}, {
		name:   "windowing out",
		values: []float64{5, 6, 5, 5, 5, 5, 5},
		expect: 5,
	}, {
		name:   "windowing out with gaps",
		values: []float64{6, 5, 2, 1},
		indexFunc: func(i int) int {
			if i >= 3 {
				return i + 3
			}

			return i
		},
		expect: 2,
	}, {
		name:   "windowing out 2",
		values: []float64{5, 6, 5, 7, 5, 5, 1},
		expect: 7,
	}, {
		name:   "windowing out 3",
		values: []float64{5, 8, 5, 7, 5, 5},
		expect: 8,
	}, {
		name:   "windowing out 4",
		values: []float64{5, 8, 5, 7, 5, 5, 1},
		expect: 7,
	}, {
		name:   "windowing out 5",
		values: []float64{5, 8, 5, 7, 5, 5, 1, 4, 4, 4},
		expect: 5,
	}, {
		name:   "windowing out 6",
		values: []float64{5, 8, 5, 7, 5, 5, 1, 4, 4, 4, 4},
		expect: 4,
	}, {
		name:   "windowing out 7",
		values: []float64{5, 8, 5, 7, 5, 5, 1, 4, 4, 4, 4, 9},
		expect: 9,
	}, {
		name:   "windowing out 8",
		values: []float64{5, 8, 5, 7, 5, 5, 1, 4, 4, 4, 4, 9, 3, 4, 2, 1, 0},
		expect: 4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max := window.NewMax(5)

			indexFunc := func(i int) int { return i }
			if tt.indexFunc != nil {
				indexFunc = tt.indexFunc
			}

			for i, v := range tt.values {
				max.Record(indexFunc(i), v)
			}

			if got, want := max.Current(), tt.expect; got != want {
				t.Errorf("Current() = %f, expected %f", got, want)
			}
		})
	}
}
