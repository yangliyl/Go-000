package window

import (
	"math"
	"sync"
	"time"
)

// Window struct
type Window struct {
	sync.RWMutex
	size       int
	sum        float64
	windowTime time.Duration
	bucketTime time.Duration
	buckets    []*bucket
}

type bucket struct {
	time  time.Time
	value float64
}

// NewWindow Create Window
func NewWindow(windowTime, bucketTime time.Duration) *Window {
	return &Window{
		size:       int(windowTime.Milliseconds() / bucketTime.Milliseconds()),
		windowTime: windowTime,
		bucketTime: bucketTime,
		buckets:    make([]*bucket, 0),
	}
}

// Add func
func (w *Window) Add(val float64) {
	if val == 0 {
		return
	}

	w.Lock()
	defer w.Unlock()
	b := w.getCurrentBucket()
	b.value += val
	w.sum += val
	w.removeOldBuckets()
}

// Sum func
func (w *Window) Sum() float64 {
	return w.sum
}

// Avg func
func (w *Window) Avg() float64 {
	return w.sum / float64(w.size)
}

// Max func
func (w *Window) Max() float64 {
	var max float64

	start := time.Now().Add(-w.windowTime)

	w.RLock()
	defer w.RUnlock()

	for _, b := range w.buckets {
		if b.time.After(start) && b.value > max {
			max = b.value
		}
	}

	return max
}

// Min func
func (w *Window) Min() float64 {
	min := math.MaxFloat64

	start := time.Now().Add(-w.windowTime)

	w.RLock()
	defer w.RUnlock()

	for _, b := range w.buckets {
		if b.time.After(start) && b.value < min {
			min = b.value
		}
	}

	return min
}

func (w *Window) getCurrentBucket() *bucket {
	now := time.Now()

	var b *bucket
	if len(w.buckets) == 0 {
		b = &bucket{time: now}
		w.buckets = append(w.buckets, b)
	}

	b = w.buckets[len(w.buckets)-1]
	if now.Sub(b.time) >= w.bucketTime {
		b = &bucket{time: now}
		w.buckets = append(w.buckets, b)
	}

	return b
}

func (w *Window) removeOldBuckets() {
	start := time.Now().Add(-w.windowTime)

	for _, b := range w.buckets {
		if b.time.Before(start) {
			w.buckets = w.buckets[1:]
			w.sum -= b.value
		}
	}
}
