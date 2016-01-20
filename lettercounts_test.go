package part3

import (
	"testing"
	"reflect"
)

func TestLetterCountsExamples(t *testing.T) {
	if !reflect.DeepEqual(Lettercounts("foo bar baz"), map[rune]int{'f':1, 'r':1, 'z':1,'a':2, 'b':2, 'o':2}){
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("foobarbaz"), map[rune]int{'f':1, 'r':1, 'z':1,'a':2, 'b':2, 'o':2}){
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("barfoo  baz"), map[rune]int{'f':1, 'r':1, 'z':1,'a':2, 'b':2, 'o':2}){
		t.Error("Failed.")
	}
	if !reflect.DeepEqual(Lettercounts("foobar baz"), map[rune]int{'f':1, 'r':1, 'z':1,'a':2, 'b':2, 'o':2}){
		t.Error("Failed.")
	}
}

func TestMoar(t *testing.T) {
	if !reflect.DeepEqual(Lettercounts("Matt Whisenhunt"), map[rune]int{'t':3, 'r':1, 'z':1,'a':2, 'b':2, 'o':2}){
		t.Error("Failed.")
	}
}
