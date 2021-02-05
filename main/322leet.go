package main

func main() {

}

func coinChange(coins []int, amount int) int {

	if amount < 0 {
		return 0
	}

	changeMaker := make([]int, amount+1)

	for idx := range changeMaker {
		changeMaker[idx] = amount + 1
	}

	changeMaker[0] = 0

	for i := 1; i <= len(changeMaker); i++ {
		for _, coin := range coins {
			if coin <= changeMaker[i] {
				if changeMaker[i-coin]+1 < changeMaker[i] {
					changeMaker[i] = changeMaker[i-coin] + 1
				}
			}
		}
	}

	if changeMaker[len(changeMaker)-1] > amount {
		return -1
	}

	return changeMaker[len(changeMaker)-1]
}
