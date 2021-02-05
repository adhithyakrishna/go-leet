package main

import "fmt"

func main() {
	d := isAnagram("anagram", "nagaram")
	fmt.Println(d)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	d := make([]int, 26)
	count := 0
	for i := 0; i < len(s); i++ {
		d[s[i]-'a']++
		count++
	}

	for i := 0; i < len(t); i++ {
		d[t[i]-'a']--
		if d[t[i]-'a'] < 0 {
			return false
		}
		count--
	}

	if count == 0 {
		return true
	}

	return false
}
