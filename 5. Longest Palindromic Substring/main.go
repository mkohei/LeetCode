package main

func longestPalindrome(s string) string {
	longest := ""

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]
			if isPalindrome(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	println(longestPalindrome("babad"))
}
