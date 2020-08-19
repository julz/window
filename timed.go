package window

import (
	"math"
	"time"
)

type TimedMax struct {
	max         *Max
	granularity time.Duration
}

func NewTimedMax(duration, granularity time.Duration) *TimedMax {
	buckets := int(math.Ceil(float64(duration) / float64(granularity)))
	return &TimedMax{max: NewMax(buckets), granularity: granularity}
}

func (t *TimedMax) Record(now time.Time, value float64) {
	index := int(now.Unix()) / int(t.granularity.Seconds())
	t.max.Record(index, value)
}

func (t *TimedMax) Current() float64 {
	return t.max.Current()
}
