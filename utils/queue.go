package utils

import "errors"

type queueItem[T any] struct {
	content T
	link    *queueItem[T]
}

type Queue[T any] struct {
	first *queueItem[T]

	count int
}

func (q *Queue[T]) Count() int {
	return q.count
}

func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue[T]) Enqueue(content T) {
	q.count++
	var queuedItem = &queueItem[T]{content, nil}

	if q.first == nil {
		q.first = queuedItem
		return
	}

	var node = q.first
	for node.link != nil {
		node = node.link
	}

	node.link = queuedItem
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.first == nil {
		return nil, errors.New("queue is empty")
	}

	var item = q.first.content
	q.first = q.first.link
	q.count--
	return &item, nil
}
