/*
Isomorphic Strings - Given two strings s and t, determine if they are isomorphic. Two strings s and t are isomorphic if the characters in s can be replaced to get t. All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself. s and t consist of any valid ascii character.
*/

package main

func isomorphicString(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if val, exists := sToT[charS]; exists {
			if val != charT {
				return false
			}
		} else {
			sToT[charS] = charT
		}

		if val, exists := tToS[charT]; exists {
			if val != charS {
				return false
			}
		} else {
			tToS[charT] = charS
		}
	}

	return true
}
