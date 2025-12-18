/*

Coding Exercise: Monotonic Array
Question

An array is monotonic if it is either monotone increasing or monotone decreasing. An array is monotone increasing if all its elements from left to right are non-decreasing. An array is monotone decreasing if all  its elements from left to right are non-increasing. Given an integer array return true if the given array is monotonic, or false otherwise.

*/

const isMonotonic = (arr:number[]) => {
    let inc:boolean = true;
    let dec:boolean = true;

    for ( let i = 0; i < arr.length - 1; i ++) {
        if ( arr[i] > arr[i+1]){
            inc = false
        }else if (arr[i] < arr[i+1]) {
            dec = false
        }
    }
    return inc || dec
}