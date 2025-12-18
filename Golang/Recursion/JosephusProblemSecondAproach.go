/*
	Coding Exercise: Josephus problem
	There are n friends that are playing a game. The friends are sitting in a circle and are numbered from 1 to n in clockwise order.
	More formally, moving clockwise from the ith friend brings you to the (i+1)th friend for 1 <= i < n, and moving clockwise from the nth friend brings you to the 1st friend.
	The rules of the game are as follows: Start at the 1st friend. Count the next k friends in the clockwise direction including the friend you started at.
	The counting wraps around the circle and may count some friends more than once.
	The last friend you counted leaves the circle and loses the game.
	If there is still more than one friend in the circle, go back to step 2 starting from the friend immediately clockwise of the friend who just lost and repeat.
	Else, the last friend in the circle wins the game. Given the number of friends, n, and an integer k, return the winner of the game.
*/

// in this approach we will find the winner with n people and k step count by using the winner of n-1 people and k step count
// for example if we have 5 people and k=2 then the winner can be found by finding the winner of 4 people and k=2
// 5 people: 1 2 3 4 5, and eliminate 2 -> 1 3 4 5 then the winner of 4 people with k=2 is 1 so we can say that the winner of 5 people is 3 because 3 is the next person after 2

package main

func findTheWinner2(n int, k int) int {
	var helper func(int) int

	helper = func(n int) int {
		if n == 1 {
			return 0
		}
		return (helper(n-1) + k) % n
	}

	return helper(n) + 1
}
