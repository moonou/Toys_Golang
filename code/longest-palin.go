package code

import "fmt"

func Longest_palin() {
	fmt.Println(longestPalindrome("aaaabaaa"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	min, max := 0, 0
	for i := range s {
		for j := range s[i:] {
			if (i+j+1)-i < max-min {
				continue
			}
			if strtest(s[i : i+j+1]) {
				min = i
				max = i + j + 1
			}
		}
	}
	return s[min:max]
}

func strtest(s string) bool {
	middle := len(s) / 2
	if len(s) == 0 {
		return false
	}
	for i, _ := range s[:middle] {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
