package window

import (
	"math"
	"time"
)

// TimedMax is a descending minima window whose indexes are calculated based on
// time.Time values.
type TimedMax struct {
	max         *Max
	granularity time.Duration
}

// NewTimedMax creates a new TimedMax window.
func NewTimedMax(duration, granularity time.Duration) *TimedMax {
	buckets := int(math.Ceil(float64(duration) / float64(granularity)))
	return &TimedMax{max: NewMax(buckets), granularity: granularity}
}

// Record records a value in the bucket derived from the given time.
func (t *TimedMax) Record(now time.Time, value float64) {
	index := int(now.Unix()) / int(t.granularity.Seconds())
	t.max.Record(index, value)
}

// Current returns the current maximum value observed in the previous window
// duration.
func (t *TimedMax) Current() float64 {
	return t.max.Current()
}
