package anagram

// func IsAnagram(src, tgt string) bool {
// 	if len(src) != len(tgt) {
// 		return false
// 	}

// 	srcArr := []rune(src)
// 	sort.Slice(srcArr, func(i, j int) bool {
// 		return srcArr[i] < srcArr[j]
// 	})

// 	tgtArr := []rune(tgt)
// 	sort.Slice(tgtArr, func(i, j int) bool {
// 		return tgtArr[i] < tgtArr[j]
// 	})

// 	for i := 0; i < len(srcArr); i++ {
// 		if srcArr[i] != tgtArr[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func IsAnagram(src, tgt string) bool {
// 	if len(src) != len(tgt) {
// 		return false
// 	}

// 	srcArr := []byte(src)
// 	sort.Slice(srcArr, func(i, j int) bool {
// 		return srcArr[i] < srcArr[j]
// 	})

// 	tgtArr := []byte(tgt)
// 	sort.Slice(tgtArr, func(i, j int) bool {
// 		return tgtArr[i] < tgtArr[j]
// 	})

// 	return bytes.Equal(srcArr, tgtArr)
// }

func IsAnagram(src, tgt string) bool {
	if len(src) != len(tgt) {
		return false
	}

	srcMap := make(map[rune]int)
	for _, letter := range src {
		srcMap[letter]++
	}

	tgtMap := make(map[rune]int)
	for _, letter := range tgt {
		tgtMap[letter]++
	}

	for letter, srcCount := range srcMap {
		if tgtCount, ok := tgtMap[letter]; !ok || srcCount != tgtCount {
			return false
		}
	}

	return true
}
