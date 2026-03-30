/*

Minimum Window Substring:

Given two strings s and t of lengths m and n respectively, return the minimum window

substring such that every character int(including duplicates) is included in the window. If there is no such substring, returnthe empty string"".



The testcases will be generated such that the answer is unique.
*/

package main

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	countT := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	have := 0
	need := len(countT)

	resLeft, resRight := -1, -1
	resLen := int(^uint(0) >> 1) // max int

	left := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if val, ok := countT[c]; ok && window[c] == val {
			have++
		}

		for have == need {
			if (right - left + 1) < resLen {
				resLeft = left
				resRight = right
				resLen = right - left + 1
			}

			window[s[left]]--
			if val, ok := countT[s[left]]; ok && window[s[left]] < val {
				have--
			}

			left++
		}
	}

	if resLen == int(^uint(0)>>1) {
		return ""
	}

	return s[resLeft : resRight+1]
}
