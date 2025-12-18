/*
Coding Exercise: Sorted Squared Array
Question

You are given an array of Integers in which each subsequent value is not less than the previous value. Write a function that takes this array as an input and returns a new array with the squares of each number sorted in ascending order.

Remember

You can solve this question in multiple ways.

Think about the following -

1.What would be the brute force way of solving this question ? What would be the Time and Space complexity of this approach?

2.Is there a better way to solve this with a more optimum Time complexity than the Brute force way ?

*/

const sortSquare = (arr:number[]):number[] => {
    let result:number[] = new Array(arr.length)
    let i = 0
    let j = arr.length  - 1
    let k = arr.length  - 1

    while (i <= j ) {
        if (arr[i]**2 > arr[j]**2) {
            result[k] = arr[i]**2
            i++
        } else {
            result[k] = arr[j]**2
            j--
        }
        k--
    }
    return result
}

console.log(sortSquare([-9, -5, 2, 3, 4, 10]))
