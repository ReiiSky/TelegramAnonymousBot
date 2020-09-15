package module

import (
	"container/list"
	"sync"
)

// SimpleQueue ..
type SimpleQueue struct {
	*list.List
	sync.Mutex
}

// Insert ..
func (q *SimpleQueue) Insert(v interface{}) {
	q.Lock()
	q.PushBack(v)
	q.Unlock()
}

// Take ..
func (q *SimpleQueue) Take() interface{} {
	q.Lock()
	f := q.Front()
	q.Remove(f)
	q.Unlock()
	return f.Value
}
