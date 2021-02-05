package main

import "fmt"

func main() {
	data := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(data, k)

	for _, v := range result {
		fmt.Printf("%d ", v)
	}

}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	bucket := make([][]int, len(nums))
	var res []int
	for idx, ele := range nums {
		countMap[ele]++
		bucket[idx] = []int{}
	}

	for key, val := range countMap {
		bucket[val] = append(bucket[val], key)
	}

	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		val := bucket[i]

		for _, v := range val {
			if k > 0 && len(val) > 0 {
				res = append(res, v)
				k = k - 1
			} else {
				break
			}
		}
	}
	return res
}
