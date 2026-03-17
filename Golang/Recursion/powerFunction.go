/*
Power function

Compute xⁿ recursively

Example: power(2, 5) → 32

*/

package main

func powerFunc(base int, exponent int) int {

	if exponent == 0 {
		return 1
	}

	return powerFunc(base, exponent-1) * base
}
