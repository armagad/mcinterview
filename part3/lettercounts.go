package part3

import (
	"strings"
	"unicode/utf8"
	"sort"
)

/*
*/

type Lettercounts struct {
	letters map[rune]uint
	length  int
}

func Countletters(word string) (l Lettercounts) {
	l.letters = make(map[rune]uint)
	for _, c := range strings.ToLower(word) {

		if c != ' ' {

			l.letters[c] += 1
			if l.letters[c] == 1 {
				l.length += utf8.RuneLen(c)
			}
		}
	}

	return l
}

func (l Lettercounts) String() string {
	str := make([]byte, l.length)
	pos := 0
	l.Foreach(func(v uint, k rune) {
		pos += utf8.EncodeRune(str[pos:], k)
	})
	return string(str)
}

func (l Lettercounts) Foreach(fn func(uint, rune)) {
	pairs := make([]kv, len(l.letters))

	i := 0
	for k, v := range l.letters {
		pairs[i] = kv{k, v}
		i++
	}
	sort.Sort(byCountAlpha(pairs))

	for _, pair := range pairs {
		fn(pair.v, pair.k)
	}
}

// Key/value pair and sort.Interface implementation

type kv struct {
	k rune
	v uint
}

type byCountAlpha []kv

func (a byCountAlpha) Len() int           { return len(a) }
func (a byCountAlpha) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCountAlpha) Less(i, j int) bool {
	if a[i].v == a[j].v {
		return a[i].k < a[j].k
	}
	return a[i].v > a[j].v // Greater than to sort DESC
}
