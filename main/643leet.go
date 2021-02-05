package main

import "fmt"

func main() {
	d := countSubstrings("abc")
	fmt.Print("result", d)
}

func countSubstrings(s string) int {

	if len(s) <= 0 {
		return 0
	}

	var result = make([]int, 1, 1)

	for i := 0; i < len(s)-1; i++ {

		expandAroundCenter(i, i, s, result)
		expandAroundCenter(i, (i + 1), s, result)
	}

	return result[0]
}

func expandAroundCenter(left int, right int, s string, result []int) {
	for left >= 0 && right < len(s) {

		if s[left] == s[right] {
			result[0] = result[0] + 1
			left = left - 1
			right = right + 1
		} else {
			return
		}
	}
}
