/*
Coding Exercise(Tower of Hanoi)
The tower of Hanoi is a famous puzzle where we have three rods and N disks. The objective of the puzzle is to move the entire stack to another rod.
You are given the number of discs N. Initially, these discs are in the rod 1.
You need to print all the steps of discs movement so that all the discs reach the 3rd rod. Also, you need to find the total moves.



Note: The discs are arranged such that the top disc is numbered 1 and the bottom-most disc is numbered N.
Also, all the discs have different sizes and a bigger disc cannot be put on the top of a smaller disc.
Refer the provided link to get a better clarity about the puzzle.



EXAMPLE:
For Input (N = 2)
Output:
move disk 1 from rod 1 to rod 2
move disk 2 from rod 1 to rod 3
move disk 1 from rod 2 to rod 3
3
*/

package main

import "fmt"

func toh(n int, from string, to string, aux string) int {
	counter := 0
	var helper func(n int, from string, to string, aux string)

	helper = func(n int, from string, to string, aux string) {
		if n <= 0 {
			return
		}
		// base case
		if n == 1 {
			fmt.Printf("move disk %d from %s to %s\n", n, from, to)
			counter++
			return
		}

		// recursive case
		helper(n-1, from, aux, to)
		fmt.Printf("move disk %d from %s to %s\n", n, from, to)
		counter++
		helper(n-1, aux, to, from)
	}

	helper(n, from, to, aux)
	return counter
}
