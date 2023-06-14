package _25_valid_palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := regexp.MustCompile("[^a-z0-9]")
	s = r.ReplaceAllString(s, "")
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
