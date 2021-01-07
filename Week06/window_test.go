package window

import (
	"fmt"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	testCase := []struct {
		windowTime time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{0.5, 1.5, 2.5, 3.5, 4.5}, 12.5},
		{5 * time.Second, 500 * time.Millisecond, []float64{1, 2, 3, 4, 5}, 15},
	}

	for _, c := range testCase {
		t.Run(fmt.Sprintf("windowTime: %+v, bucketTime: %+v", c.windowTime, c.bucketTime), func(t *testing.T) {
			w := NewWindow(c.windowTime, c.bucketTime)
			for _, x := range c.data {
				w.Add(x)
				time.Sleep(c.bucketTime)
			}

			if got := w.Sum(); got != c.want {
				t.Errorf("got %f, want %f", got, c.want)
			}
		})
	}
}

func TestAvg(t *testing.T) {
	testCase := []struct {
		windowTime time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{0.5, 1.5, 2.5, 3.5, 4.5}, 1.25},
		{5 * time.Second, 500 * time.Millisecond, []float64{1, 2, 3, 4, 5}, 1.5},
	}

	for _, c := range testCase {
		t.Run(fmt.Sprintf("windowTime: %+v, bucketTime: %+v", c.windowTime, c.bucketTime), func(t *testing.T) {
			w := NewWindow(c.windowTime, c.bucketTime)
			for _, x := range c.data {
				w.Add(x)
				time.Sleep(c.bucketTime)
			}

			if got := w.Avg(); got != c.want {
				t.Errorf("got %f, want %f", got, c.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	testCase := []struct {
		windowTime time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{10, 11, 9, 8, 12, 7, 6}, 12},
	}

	for _, c := range testCase {
		t.Run(fmt.Sprintf("windowTime: %+v, bucketTime: %+v", c.windowTime, c.bucketTime), func(t *testing.T) {
			w := NewWindow(c.windowTime, c.bucketTime)
			for _, x := range c.data {
				w.Add(x)
				time.Sleep(1 * time.Second)
			}

			if got := w.Max(); got != c.want {
				t.Errorf("got %f, want %f", got, c.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	testCase := []struct {
		windowTime time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{10, 11, 9, 8, 12, 7, 6}, 6},
	}

	for _, c := range testCase {
		t.Run(fmt.Sprintf("windowTime: %+v, bucketTime: %+v", c.windowTime, c.bucketTime), func(t *testing.T) {
			w := NewWindow(c.windowTime, c.bucketTime)
			for _, x := range c.data {
				w.Add(x)
				time.Sleep(1 * time.Second)
			}

			if got := w.Min(); got != c.want {
				t.Errorf("got %f, want %f", got, c.want)
			}
		})
	}
}
