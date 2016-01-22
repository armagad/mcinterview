package part3

func Lettercounts(word string) (letters map[rune]int) {
	letters = make(map[rune]int)
	for _, c := range word{
		if c != ' ' {
			letters[c] += 1
		}
	}

	return letters
}
