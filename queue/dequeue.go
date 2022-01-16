package queue

import (
	"errors"
	"sync"
)

var (
	ErrDeQueueIsEmpty = errors.New("queue is empty")
	ErrNilItem        = errors.New("should not add a nil item to queue")
)

// DeQueue, implemented with slice, is a double-ended queue
// which supports adding and removing items
// from either the front or the back of the data structure
type DeQueue struct {
	elements []interface{}

	mu sync.Mutex
}

// return an empty double end queue
func New() *DeQueue {
	return &DeQueue{
		elements: make([]interface{}, 0),
	}
}

// check if the queue is empty
func (dq *DeQueue) IsEmpty() bool {
	return len(dq.elements) == 0
}

func (dq *DeQueue) Size() int {
	return len(dq.elements)
}

// instead of dq.elements = append([]interface{}{item}, dq.elements...)
// is more efficient for this method will not allocate new memory
// but as size of the queue becomes larger and larger, addfirst becomes slower, so it is inappropriate to do benchmark test
func (dq *DeQueue) AddFirst(item interface{}) error {
	if item == nil {
		return ErrNilItem
	}
	dq.mu.Lock()
	defer dq.mu.Unlock()

	dq.elements = append(dq.elements, 0)
	copy(dq.elements[1:], dq.elements)
	dq.elements[0] = item
	return nil
}

// this method will allocate new memory
// do not use
// func (dq *DeQueue) AddFirstSlow(item interface{}) {
// 	dq.mu.Lock()
// 	defer dq.mu.Unlock()

// 	dq.elements = append([]interface{}{item}, dq.elements...)
// }

func (dq *DeQueue) AddLast(item interface{}) error {
	if item == nil {
		return ErrNilItem
	}
	dq.mu.Lock()
	defer dq.mu.Unlock()
	dq.elements = append(dq.elements, item)
	return nil
}

func (dq *DeQueue) RemoveFirst() (interface{}, error) {
	if dq.Size() == 0 {
		return nil, ErrDeQueueIsEmpty
	}
	dq.mu.Lock()
	defer dq.mu.Unlock()
	v := dq.elements[0]
	dq.elements = dq.elements[1:]
	return v, nil
}

func (dq *DeQueue) RemoveLast() (interface{}, error) {
	if dq.Size() == 0 {
		return nil, ErrDeQueueIsEmpty
	}
	dq.mu.Lock()
	defer dq.mu.Unlock()
	v := dq.elements[dq.Size()-1]
	dq.elements = dq.elements[:dq.Size()-1]
	return v, nil
}

func (dq *DeQueue) Iterator() ([]interface{}, error) {
	if dq.Size() == 0 {
		return nil, ErrDeQueueIsEmpty
	}
	return dq.elements, nil
}
