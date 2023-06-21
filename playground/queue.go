package main

type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
