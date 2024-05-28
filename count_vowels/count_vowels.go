package countvowels

func CountVowels(str string) int {
	count := 0

	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	return count
}
