package util

import "strings"

func IsPalindrome(text string) bool {
	cleanedText := strings.ReplaceAll(strings.ToLower(text), " ", "")

	runes := []rune(cleanedText)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedText := string(runes)

	return cleanedText == reversedText
}
