package main

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() Queue {
	q := Queue{}
	return q
}

func (q *Queue) insert(data int) {
	q.data = append(q.data, data)
}

func (q *Queue) popFirst() int {
	if len(q.data) == 0 {
		return -1
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) pollLast() int {
	if len(q.data) == 0 {
		return -1
	}
	res := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	return res
}

func main() {

	data := NewQueue()

	data.insert(1)
	data.insert(2)
	data.insert(3)
	data.insert(4)

	valFirst := data.popFirst()
	valLast := data.pollLast()

	fmt.Println(valFirst)
	fmt.Println(valLast)

	fmt.Println(data)
}
