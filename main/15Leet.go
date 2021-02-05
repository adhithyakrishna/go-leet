package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		//skip if same, because the result has to be unique
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		currentVal := nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			total := currentVal + nums[left] + nums[right]
			if total == 0 {
				values := []int{nums[left], nums[right]}
				res = append(res, values)

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
