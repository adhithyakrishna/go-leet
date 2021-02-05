package main

import "fmt"

type Stack struct {
	data [][]int
}

func NewQueue() Stack {
	q := Stack{}
	return q
}

func (q *Stack) insert(data []int) {
	q.data = append([][]int{data}, q.data...)
}

func (q *Stack) pop() []int {
	if len(q.data) == 0 {
		return []int{}
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func main() {

	data := NewQueue()

	data.insert([]int{1, 1})
	data.insert([]int{1, 2})
	data.insert([]int{1, 3})
	data.insert([]int{1, 4})
	data.insert([]int{1, 5})

	valFirst := data.pop()

	fmt.Println(valFirst)
	fmt.Println(data)
}
