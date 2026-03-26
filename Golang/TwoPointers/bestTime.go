/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example:

Input: prices = [9,1,5,3,7,5]

Output: 6

Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 7), profit = 7-1 = 6.
*/

package main

func maxProfit(prices []int) int {
	left := 0
	profit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else {
			currentProfit := prices[right] - prices[left]
			if currentProfit > profit {
				profit = currentProfit
			}
		}
	}

	return profit
}
