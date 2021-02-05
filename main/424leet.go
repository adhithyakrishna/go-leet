package main

import "fmt"

func main() {
	d := characterReplacement("AABABBA", 1)
	fmt.Println(d)
}

func characterReplacement(s string, maxAllowed int) int {
	charFreq := make(map[byte]int)

	start := 0
	end := 0

	maxFrequency := 0
	var maxChar byte
	maxChar = '~'

	max := 0
	maxchanges := 0

	for end < len(s) {
		currentChar := s[end]
		val, ok := charFreq[currentChar]
		if ok {
			charFreq[currentChar] = val + 1
		} else {
			charFreq[currentChar] = 1
		}

		if charFreq[currentChar] > maxFrequency {
			maxChar = currentChar
			maxFrequency, _ = charFreq[currentChar]
		}

		maxchanges = (end - start + 1) - maxFrequency

		if maxchanges > maxAllowed {
			startChar := s[start]
			if maxChar == startChar {
				maxFrequency = maxFrequency - 1
			}

			charFreq[startChar] = charFreq[s[start]] + 1
			start++
		}

		if end-start+1 > max {
			max = end - start + 1
		}
		end++
	}

	return max
}
