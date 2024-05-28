package palindrome

func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if str[0] != str[len(str)-1] {
		return false
	}

	return IsPalindrome(str[1 : len(str)-1])
}

// func IsPalindrome(str string) bool {
// 	for i := 0; i < len(str)/2; i++ {
// 		if str[i] != str[len(str)-i-1] {
// 			return false
// 		}
// 	}

// 	return true
// }
