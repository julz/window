package window

type entry struct {
	value float64
	index int
}

// Max is a circular buffer which keeps track of the maximum value observed in a particular time.
// Based on the "ascending minima algorithm" (http://web.archive.org/web/20120805114719/http://home.tiac.net/~cri/2001/slidingmin.html).
type Max struct {
	maxima        []entry
	first, length int
}

func NewMax(size int) *Max {
	return &Max{
		maxima: make([]entry, size),
	}
}

// Record records a value for a monotonically increasing index.
func (m *Max) Record(index int, v float64) {
	// Step One: Remove any elements where v > element.
	// An element that's lower than the new element can never influence the
	// maximum again, because the current element is both larger _and_ more
	// recent.
	length := m.length
	for i := 0; i < length; i++ {
		// Search backwards because that way we can delete by just decrementing length.
		// The elements are in descending order as described in Step Three.
		if v >= m.maxima[m.index(m.first+length-i-1)].value {
			m.length--
		}
	}

	// Step Two: Remove first element if it's out of date now.
	// We only ever add at end of list, so the indexes are in ascending order.
	for index-m.maxima[m.first].index >= len(m.maxima) {
		m.length--
		m.first++

		// circle around the buffer if neccessary.
		if m.first == len(m.maxima) {
			m.first = 0
		}
	}

	// Step Three: Add the new value to the end (which maintains sorted order
	// since we removed any lesser values above, so value we're appending is
	// always smallest value in list).
	m.maxima[m.index(m.first+m.length)] = entry{value: v, index: index}
	m.length++

	// We removed an item from the list in Step Two if it was added more than
	// len(maxima) ago, so length can never be larger than len(maxima).
	if m.length > len(m.maxima) {
		panic("length exceeded buffer size. this is impossible, you win a prize.")
	}
}

// Current returns the current maximum value observed.
func (m *Max) Current() float64 {
	return m.maxima[m.first].value
}

func (m *Max) index(i int) int {
	return i % len(m.maxima)
}
