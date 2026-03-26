/*
Isomorphic Strings - Given two strings s and t, determine if they are isomorphic. Two strings s and t are isomorphic if the characters in s can be replaced to get t. All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself. s and t consist of any valid ascii character.
*/

function isIsomorphicStrings(s: string, t: string): boolean {
    if (s.length !== t.length) return false;

    const sMap: Record<string, string> = {};
    const tMap: Record<string, string> = {};

    for (let i = 0; i < s.length; i++) {
        const charS = s[i];
        const charT = t[i];

        if (!(charS in sMap)) {
            sMap[charS] = charT;
        }

        if (!(charT in tMap)) {
            tMap[charT] = charS;
        }

        if (sMap[charS] !== charT || tMap[charT] !== charS) {
            return false;
        }
    }

    return true;
}