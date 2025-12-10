/*

Given a string s, print every subsequence (not substring).
A subsequence chooses to include or exclude each character.

Example:
s = "abc"
Output:

""
"a"
"b"
"c"
"ab"
"ac"
"bc"
"abc"

*/

package main

func subsequencesRet(i int, cur string, str string, res *[]string) {
	if i == len(str) {
		*res = append(*res, cur)
		return
	}
	subsequencesRet(i+1, cur, str, res)
	subsequencesRet(i+1, cur+string(str[i]), str, res)
}

func Subsequences(str string) []string {
	var res []string
	subsequencesRet(0, "", str, &res)
	return res
}
