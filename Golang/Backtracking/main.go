package main

import "fmt"

func main() {
	nums1 := []int{1, 4}
	fmt.Println("Permutations of", nums1, "are:", permute(nums1))

	nums2 := []int{1, 4, 5}
	fmt.Println("Permutations of", nums2, "are:", permute(nums2))
}
