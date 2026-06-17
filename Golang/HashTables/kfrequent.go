package main

func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)

    // count frequency
    for _, num := range nums {
        freq[num]++
    }

    // bucket[i] chứa các số xuất hiện i lần
    bucket := make([][]int, len(nums)+1)

    for num, count := range freq {
        bucket[count] = append(bucket[count], num)
    }

    result := []int{}

    // duyệt từ frequency lớn nhất xuống
    for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
        if len(bucket[i]) > 0 {
            result = append(result, bucket[i]...)
        }
    }

    return result[:k]
}