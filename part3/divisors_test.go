package part3

import (
	"testing"
)

func TestDivisorsExamples(t *testing.T) {
	if !compareSlices(Divisors(12), []uint{2, 3, 4, 6}) {
		t.Error("12 Failed.")
	}
	if !compareSlices(Divisors(13), []uint{}) {
		t.Error("13 Failed.")
	}
	if !compareSlices(Divisors(25), []uint{5}) {
		t.Error("25 Failed.")
	}

	if compareSlices(Divisors(12), []uint{}) {
		t.Error("12 Failed.")
	}
	if compareSlices(Divisors(13), []uint{8, 9, 10, 11, 12}) {
		t.Error("13 Failed.")
	}
	if compareSlices(Divisors(25), []uint{2, 3, 4, 5}) {
		t.Error("25 Failed.")
	}
}

func TestDivisorsSamples(t *testing.T) {

	if !compareSlices(Divisors(100), []uint{2, 4, 5, 10, 20, 25, 50}) {
		t.Error("100 Failed.")
	}

	if !compareSlices(Divisors(238), []uint{2, 7, 14, 17, 34, 119}) {
		t.Error("238 Failed.")
	}
	if !compareSlices(Divisors(239), []uint{}) {
		t.Error("239 Failed.")
	}
	if !compareSlices(Divisors(240), []uint{2, 3, 4, 5, 6, 8, 10, 12, 15, 16, 20, 24, 30, 40, 48, 60, 80, 120}) {
		t.Error("240 Failed.")
	}

	if !compareSlices(Divisors(658), []uint{2, 7, 14, 47, 94, 329}) {
		t.Error("658 Failed.")
	}
	if !compareSlices(Divisors(659), []uint{}) {
		t.Error("659 Failed.")
	}
	if !compareSlices(Divisors(660), []uint{2, 3, 4, 5, 6, 10, 11, 12, 15, 20, 22, 30, 33, 44, 55, 60, 66, 110, 132, 165, 220, 330}) {
		t.Error("660 Failed.")
	}

	if !compareSlices(Divisors(998), []uint{2, 499}) {
		t.Error("998 Failed.")
	}
	if !compareSlices(Divisors(999), []uint{3, 9, 27, 37, 111, 333}) {
		t.Error("999 Failed.")
	}
	if !compareSlices(Divisors(1000), []uint{2, 4, 5, 8, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500}) {
		t.Error("1000 Failed.")
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
