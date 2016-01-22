package part3

import (
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*

I ended having to make a trade off: Provide a data structure that could be indexed like
a built in map or keep track of rune lengths at insert and alocate memory for output at O(1)

I added Json() as an example of spending O(n) to calculate the size of memory to allocate
for the output string. Even it had another trade off loop through once and append/allocate as
I went along or O(n) twice and just do one allocation. I don't know the right answer.

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

func (l Lettercounts) Get(r rune) uint { return l.letters[r] }

func (l Lettercounts) String() string {
	str := make([]byte, l.length)
	pos := 0
	l.Foreach(func(v uint, k rune) {
		pos += utf8.EncodeRune(str[pos:], k)
	})
	return string(str)
}

func (l Lettercounts) Json() string {

	length := 1
	for k, v := range l.letters {
		length += 4 + utf8.RuneLen(k) + numDigits(v)
	}
	str := make([]byte, length)
	pos := 1
	str[0] = '{'
	str[length-1] = '}'
	comma := false
	l.Foreach(func(v uint, k rune) {
		if comma {
			str[pos] = ','
			pos += 1
		} else {
			comma = true
		}

		str[pos] = '"'
		pos += 1
		pos += utf8.EncodeRune(str[pos:], k)
		str[pos] = '"'
		pos += 1
		str[pos] = ':'
		pos += 1
		digits := strconv.Itoa(int(v))
		for i := 0; i < len(digits); i++ {
			str[pos] = digits[i]
			pos += 1
		}
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

func (a byCountAlpha) Len() int      { return len(a) }
func (a byCountAlpha) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byCountAlpha) Less(i, j int) bool {
	if a[i].v == a[j].v {
		return a[i].k < a[j].k
	}
	return a[i].v > a[j].v // Greater than to sort DESC
}

func numDigits(x uint) int {
	if x < 10 {
		return 1
	}
	if x < 100 {
		return 2
	}
	if x < 1000 {
		return 3
	}
	if x < 10000 {
		return 4
	}
	if x < 100000 {
		return 5
	}
	if x < 1000000 {
		return 6
	}
	return 21
}
