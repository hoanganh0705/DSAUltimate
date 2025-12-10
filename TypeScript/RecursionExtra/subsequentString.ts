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

const subsequences = (s: string): string[] => {
    const result: string[] = [];

    const generateSubsequences = (index: number, current: string) => {
        if (index === s.length) {
            result.push(current);
            return;
        }

        // Include the current character
        generateSubsequences(index + 1, current + s[index]);

        // Exclude the current character
        generateSubsequences(index + 1, current);
    };

    generateSubsequences(0, "");
    return result;
}