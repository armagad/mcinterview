package part3

import (
	"testing"
)

func TestDivisorsExamples(t *testing.T) {
	if !compareSlices(Divisors(12), []uint{2, 3, 4, 6}) {
		t.Error("12 Failed.")
	}
	if Divisors(13) != nil {
		t.Error("13 Failed.")
	}
	if !compareSlices(Divisors(25), []uint{5}) {
		t.Error("25 Failed.")
	}
}

func compareSlices(a, b []uint) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
