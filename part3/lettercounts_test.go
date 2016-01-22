package part3

import (
	"reflect"
	"testing"
)

func TestLetterCountsExamples(t *testing.T) {
	if !reflect.DeepEqual(Lettercounts("foo bar baz"), map[rune]int{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("foobarbaz"), map[rune]int{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("barfoo  baz"), map[rune]int{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("foobar baz"), map[rune]int{'f': 1, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}


	if reflect.DeepEqual(Lettercounts("foo bar baz"), nil) {
		t.Error("Failed.")
	}
	if reflect.DeepEqual(Lettercounts("foo bar baz"), map[rune]int{'f': 1, 'o': 2}) {
		t.Error("Failed.")
	}
}

func TestMoar(t *testing.T) {
	if reflect.DeepEqual(Lettercounts("Matt Whisenhunt"), map[rune]int{'t': 3, 'r': 1, 'z': 1, 'a': 2, 'b': 2, 'o': 2}) {
		t.Error("Failed.")
	}
	if reflect.DeepEqual(Lettercounts("Matt Whisenhunt"), map[rune]int{'t':3,'h':2,'n':2,'a':1,'e':1,'i':1,'m':1,'s':1,'u':1,'w':1}) {
		t.Error("Failed.")
	}
}
