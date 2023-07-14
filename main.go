package main

import "fmt"

type ConcurrentQueue struct {
	queue []int32
}

func (q *ConcurrentQueue) Dequeue() int32 {
	if len(q.queue) == 0 {
		panic("cannot dequeue from the empty queue")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	return item
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Size() int {

	return len(q.queue)
}

func main() {

	q1 := ConcurrentQueue{
		queue: make([]int32, 0),
	}
	q1.Enqueue(1)

	q1.Enqueue(2)

	q1.Enqueue(3)
	fmt.Println(q1.Dequeue())
	fmt.Println(q1.Dequeue())
	fmt.Println(q1.Dequeue())
	fmt.Println(q1.Dequeue())
}
