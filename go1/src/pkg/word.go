package pkg

import "unicode"

func IsPalindrome(s string) bool {
	//var letters []rune
	letters := make([]rune, 0, len(s)) //优化2
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	//for i := range letters {
	//	if letters[i] != letters[len(letters)-1-i] {
	//		return false
	//	}
	//}
	//优化1
	for i := 0; i <= len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
