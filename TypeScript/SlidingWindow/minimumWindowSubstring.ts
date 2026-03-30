/*

Minimum Window Substring: 

Given two strings s and t of lengths m and n respectively, return the minimum window

substring such that every character int(including duplicates) is included in the window. If there is no such substring, returnthe empty string"".



The testcases will be generated such that the answer is unique.
 */

function minWindow(s: string, t: string): string {
    if (t === "") return "";

    const countT: Record<string, number> = {};
    const window: Record<string, number> = {};

    for (const c of t) {
        countT[c] = (countT[c] || 0) + 1;
    }

    let have = 0;
    const need = Object.keys(countT).length;

    let res: [number, number] = [-1, -1];
    let resLen = Infinity;

    let left = 0;

    for (let right = 0; right < s.length; right++) {
        const c = s[right];
        window[c] = (window[c] || 0) + 1;

        if (c in countT && window[c] === countT[c]) {
            have++;
        }

        while (have === need) {
            if ((right - left + 1) < resLen) {
                res = [left, right];
                resLen = right - left + 1;
            }

            window[s[left]]--;
            if (s[left] in countT && window[s[left]] < countT[s[left]]) {
                have--;
            }

            left++;
        }
    }

    const [l, r] = res;
    return resLen !== Infinity ? s.slice(l, r + 1) : "";
}