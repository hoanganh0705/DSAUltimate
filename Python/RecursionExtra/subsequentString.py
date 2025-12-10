'''

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

'''

def subsequents(i, cur, s):
    if (i == len(s)):
        print(cur)
        return

    subsequents(i+1, cur, s)
    subsequents(i+1, cur + s[i], s)

subsequents(0, "", "abc")

