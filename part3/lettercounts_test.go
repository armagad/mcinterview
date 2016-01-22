package part3

import (
	"fmt"
	"testing"
)

func TestCountlettersExamples(t *testing.T) {
	var l Lettercounts
	const example = "{\"a\":2,\"b\":2,\"o\":2,\"f\":1,\"r\":1,\"z\":1}"

	l = Countletters("foo bar baz")
	if l.Json() != example || !checkCounts(l, map[rune]uint{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}

	l = Countletters("foobarbaz")
	if l.Json() != example || !checkCounts(l, map[rune]uint{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}

	l = Countletters("barfoo  baz")
	if l.Json() != example || !checkCounts(l, map[rune]uint{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}

	l = Countletters("foobar baz")
	if fmt.Sprintf("%s", l) != "abofrz" || !checkCounts(l, map[rune]uint{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}

	// il.Json() != example ||  	!checkCounts(l, map[rune]uint{'f': 1, 'o': 2}) {
	// 	t.Error("Failed.")
	// }
}

// func TestMoar(t *testing.T) {
// 	if  Countletters("Matt Whisenhunt"), map[rune]int{'t': 3, 'h': 2, 'n': 2, 'a': 1, 'e': 1, 'i': 1, 'm': 1, 's': 1, 'u': 1, 'w': 1}) {
// 		t.Error(Countletters("Matt Whisenhunt"))
// 		t.Error(map[rune]int{'t': 3, 'h': 2, 'n': 2, 'a': 1, 'e': 1, 'i': 1, 'm': 1, 's': 1, 'u': 1, 'w': 1})
// 	}
//
// 	// false positives
// 	if Countletters("Matt Whisenhunt"), map[rune]int{'t': 3, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
// 		t.Error("Failed.")
// 	}
// 	if Countletters("Matt Whisenhunt"), map[rune]int{'t': 3, 'h': 2, 'n': 2, 'a': 1, 'e': 1, 'i': 1, 'M': 1, 's': 1, 'u': 1, 'W': 1}) {
// 		t.Error("Failed.")
// 	}
// }

func checkCounts(l Lettercounts, m map[rune]uint) bool {
	for k, v := range m {
		if l.Get(k) != v {
			return false
		}
	}
	return true
}
