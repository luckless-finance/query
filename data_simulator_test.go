package query

import (
	"math"
	"testing"
)

func TestTrig(t *testing.T) {
	n := 100
	t2 := Trig{amplitude: 10.0, periods: 1, hOffset: 20.0}
	xs := sequence(n)
	t2.sinArr(&xs)
	for i := range xs {
		if i == 0 {
			if math.Round(math.Abs(xs[i])) != 0 {
				t.Errorf("sin should be 0 at x=%d\nx_%d = %f", i, i, xs[i])
			}
		} else if i <= n/4 {
			if xs[i] < xs[i-1] {
				t.Errorf("sin should be increasing at x=%d\nx_%d = %f\nx_%d = %f", i, i, xs[i], i-1, xs[i-1])
			}
		} else if i < 3*n/4 {
			if xs[i] > xs[i-1] {
				t.Errorf("sin should be decreasing at x=%d\nx_%d = %f\nx_%d = %f", i, i, xs[i], i-1, xs[i-1])
			}
		} else if i < n-1 {
			if xs[i] < xs[i-1] {
				t.Errorf("sin should be increasing at x=%d\nx_%d = %f\nx_%d = %f", i, i, xs[i], i-1, xs[i-1])
			}
		} else {
			if math.Round(math.Abs(xs[i])) != float64(0) {
				t.Errorf("sin should be 0 at x=%d\nx_%d = %f", i, i, xs[i])
			}
		}
	}
	if len(xs) != n {
		t.Errorf(
			"expected len %d found %d",
			n,
			len(xs),
		)
	}
	if cap(xs) != n {
		t.Errorf(
			"expected len %d found %d",
			n,
			len(xs),
		)
	}
}

func TestPlot(t *testing.T) {
	plotFoo()
}
