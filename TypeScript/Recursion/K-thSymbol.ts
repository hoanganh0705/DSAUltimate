/*

Coding Exercise (k-th symbol in Grammar)
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row. Now in every subsequent row, we look at the previous row and replace each 
occurrence of 0 with 01, and each occurrence of 1 with 10. For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110. Given two 
integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

*/

const KthSymbol = (n:number, k:number ): number => {
    if (n==0)
        return 0

    let half = 2**(n-1)
    if (k <= half){
        return KthSymbol(n-1, k)
    } else {
        return 1 - KthSymbol(n-1, k-half)
    }
}

console.log(KthSymbol(3,1))
console.log(KthSymbol(3,2))
console.log(KthSymbol(3,3))
console.log(KthSymbol(3,4))