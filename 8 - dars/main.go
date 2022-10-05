package main

import (
	"fmt"
)

type queue struct {
	items []int
}

func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
}


func (q *queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) peek() int {
	return q.items[0]
}

func (q *queue) print() {
	fmt.Println(q.items)
}

func (q *queue) size() int {
	return len(q.items)
}

func works() {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.print()
	fmt.Println(q.dequeue())
	q.print()
	fmt.Println(q.peek())
	q.print()
	fmt.Println(q.size())
	q.print()
	fmt.Println(q.isEmpty())
	q.print()
}

func main() {
	works()
}
