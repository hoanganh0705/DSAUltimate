/*

Coding Exercise: Power Sum
Question:

Power Sum - Let’s define a peculiar type of array in which each element is either an integer or another peculiar array.
Assume that a peculiar array is never empty.
Write a function that will take a peculiar array as its input and find the sum of its elements.
If an array is an element in the peculiar array you have to convert it to it’s equivalent value so that you can sum it with the other elements.
Equivalent value of an array is the sum of its elements raised to the number which represents how far nested it is. For e.g. [2,3[4,1,2]] = 2+3+ (4+1+2)^2

another example - [1,2,[7,[3,4],2]] = 1 + 2 +( 7+(3+4)^3+2)^2

*/

package main

import "math"

func powersum(arr []interface{}, depth int) int {
	total := 0
	for _, v := range arr {
		switch value := v.(type) {
		case int:
			total += value
		case []interface{}:
			total += powersum(value, depth+1)
		}
	}
	return int(math.Pow(float64(total), float64(depth)))
}
