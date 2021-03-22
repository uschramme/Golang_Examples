package main

import (
	"fmt"
	"strings"
	"unicode"
)

// check if a string is a palindrome
func isPalindrome(text string) bool {
	var sb strings.Builder // a Builder is used to minimize memory copying
	// we only need letters converted to lowercas
	// no whitespaces, numbers or punctuations
	sb.Grow(len(text))
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			sb.WriteRune(unicode.ToLower(ch))
		}
	}
	text = sb.String()
	max := len(text) - 1
	// compare the first letter with the last one,
	// then the second letter with the second last
	// and so on...
	for i := 0; i < max/2; i++ {
		if text[i] != text[max-i] {
			return false // return false if the letters are not equal (no palindrome)
		}
	}
	return true // palindrome
}

func main() {
	// some tests...
	phrases := []string{"A man, a plan, a canal – Panama",
		"This is a wonderful palindrome!",
		"Madam, I'm Adam",
		"eye",
		"ear",
		"Never odd or even",
		"Do geese see God?",
		"Step on no pets",
		"Hannah",
		"Trug Tim eine so helle Hose nie mit Gurt?"} // the last one is German ;-)

	for _, test := range phrases {
		fmt.Println(test+" =", isPalindrome(test))
	}
}
