/*
Repeated DNA Sequence:

The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

•For example, "ACGAATTCCG" is a DNA sequence.

When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.

Example:

Input: s = ”GAAAATCCCCGAAAATCCCCGAAAAAGGGTTT"

Output: [”GAAAACCCCC",”TCCCCGAAAA"]
*/

package main

func findRepeatedDnaSequences(s string) []string {
	L := 10
	seen := make(map[string]bool)
	output := make(map[string]bool)

	for start := 0; start <= len(s)-L; start++ {
		temp := s[start : start+L]

		if seen[temp] {
			output[temp] = true
		}

		seen[temp] = true
	}

	result := []string{}
	for key := range output {
		result = append(result, key)
	}

	return result
}
