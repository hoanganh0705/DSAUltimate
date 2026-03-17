/*
	Count down

Print numbers from n to 1 (no loops allowed)
*/

package main

import "fmt"

func countDown(n int) {
	if n <= 0 {
		fmt.Println("Not Valid")
		return
	}

	fmt.Println(n)
	countDown(n - 1)
}
