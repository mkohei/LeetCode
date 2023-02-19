package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	maxLen := 1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			sub := s[i:j]
			if strings.Contains(sub, string(s[j])) {
				break
			} else {
				sub += string(s[j])
			}
			if len(sub) > maxLen {
				maxLen = len(sub)
			}
		}
	}
	return maxLen
}

func main() {
	println(lengthOfLongestSubstring("dvdf"))
}