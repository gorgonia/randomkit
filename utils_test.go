package randomkit

import (
	"math"
	"testing"
	"testing/quick"
)

// taken from the Go Stdlib package math
func tolerancef64(a, b, e float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}
func closeenoughf64(a, b float64) bool { return tolerancef64(a, b, 1e-8) }
func closef64(a, b float64) bool       { return tolerancef64(a, b, 1e-14) }
func veryclosef64(a, b float64) bool   { return tolerancef64(a, b, 4e-16) }
func soclosef64(a, b, e float64) bool  { return tolerancef64(a, b, e) }
func alikef64(a, b float64) bool {
	switch {
	case math.IsNaN(a) && math.IsNaN(b):
		return true
	case a == b:
		return math.Signbit(a) == math.Signbit(b)
	}
	return false
}

func TestKahan(t *testing.T) {
	t.Skip("Skipping for now until I figure out the NaN vs Inf thing")
	f := func(a []float64) bool {
		ret := Kahan(a)
		var nonacc float64
		for _, v := range a {
			nonacc += v
		}

		if !closeenoughf64(nonacc, ret) {
			t.Errorf("Ret %v, NonAcc %v ", ret, nonacc)
			return false
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}

	// a := []float64{1.079970185053218e+307, 1.5825516084036779e+308, 5.4555987234370485e+306, -9.727012296313056e+306, 8.763035333718494e+307, 6.477310198212963e+307, 7.854096166428732e+307, -9.811855267155831e+307, 1.267420017454724e+307, -9.540107655382375e+307, 1.6770056604943127e+308, -1.7953893965346735e+308, -3.7004694960120873e+307, 3.6848140163768457e+307, 6.057257631436204e+306, -1.3886066521541118e+308, -1.0626303072556577e+308, -1.346887975367412e+308, 1.3392100989275053e+308, -1.31389670035412e+308, 2.059313684532517e+306, -1.5864759969158681e+308, -6.70805377670628e+307, -4.617209452070799e+307, -1.6719743881223616e+308, 3.4539618399935834e+307, 5.904156430131649e+307, 1.1419756891135853e+308, 7.525609989511983e+307, -1.0271168192317348e+308, 4.715480231454269e+305, 1.2858956602022002e+308, -8.783449469894034e+306, 7.780522192197365e+306, -1.042476518216996e+308, 7.403935286798207e+307, 9.369162729779181e+307}
	// r := Kahan(a, -1)
	// var q float64
	// for i := 0; i < len(a); i++ {
	// 	q += a[i]
	// }
	// log.Printf("%v %v", r, q)
}
