package queue

import "container/list"

type Queue struct {
	v *list.List
}

func (q *Queue) Enqueue(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Dequeue() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}
