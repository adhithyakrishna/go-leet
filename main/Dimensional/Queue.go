package main

import "fmt"

type Queue struct {
	data [][]int
}

func NewQueue() Queue {
	q := Queue{}
	return q
}

func (q *Queue) insert(data []int) {
	q.data = append(q.data, data)
}

func (q *Queue) popFirst() []int {
	if len(q.data) == 0 {
		return []int{}
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) pollLast() []int {
	if len(q.data) == 0 {
		return []int{}
	}
	res := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return res
}

func main() {

	data := NewQueue()

	data.insert([]int{1, 1})
	data.insert([]int{1, 2})
	data.insert([]int{1, 3})
	data.insert([]int{1, 4})

	valFirst := data.popFirst()
	valLast := data.pollLast()

	fmt.Println(valFirst)
	fmt.Println(valLast)

	fmt.Println(data)
}
