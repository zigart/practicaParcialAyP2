package funciones

import (
	"errors"
)

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](longitud int) *Queue[T] {
	queue := new(Queue[T])
	queue.items = make([]T, 0, longitud)
	return queue
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (v T, err error) {
	if q.IsEmpty() {
		return v, errors.New("La cola esta vacia")
	}

	output := q.items[0]
	q.items = q.items[1:len(q.items)]
	return output, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len((*q).items) == 0
}

func (q *Queue[T]) Front() (T, error) {
	output := q.items[0]
	return output, nil
}
