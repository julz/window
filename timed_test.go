package window_test

import (
	"testing"
	"time"

	"github.com/julz/window"
)

func TestTimedWindowMax(t *testing.T) {
	type entry struct {
		time  time.Time
		value float64
	}

	now := time.Now()

	tests := []struct {
		name   string
		expect float64
		values []entry
	}{{
		name: "single value",
		values: []entry{{
			time:  now,
			value: 5,
		}},
		expect: 5,
	}, {
		name: "two values in same second",
		values: []entry{{
			time:  now,
			value: 6,
		}, {
			time:  now.Add(500 * time.Millisecond),
			value: 5,
		}},
		expect: 6,
	}, {
		name: "two values",
		values: []entry{{
			time:  now,
			value: 5,
		}, {
			time:  now.Add(1 * time.Second),
			value: 8,
		}},
		expect: 8,
	}, {
		name: "time gap",
		values: []entry{{
			time:  now,
			value: 5,
		}, {
			time:  now.Add(6 * time.Second),
			value: 4,
		}},
		expect: 4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max := window.NewTimedMax(5*time.Second, 1*time.Second)

			for _, v := range tt.values {
				max.Record(v.time, v.value)
			}

			if got, want := max.Current(), tt.expect; got != want {
				t.Errorf("Current() = %f, expected %f", got, want)
			}
		})
	}
}
