package model

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/Satssuki/tele-anon-bot-queue/internal/module"
)

// UserQueue function to wrap Queue and Cache for more functionalities
type UserQueue struct {
	Queue module.SimpleQueue
	Cache module.SafeCache
	sync.Mutex
}

func (userQ *UserQueue) delCache(userID string) {
	userQ.Cache.Delete(userID)
}

func (userQ *UserQueue) insertCache(userID string) {
	userQ.Cache.Insert(userID, "1")
}

// Len ..
func (userQ *UserQueue) Len() int {
	return userQ.Queue.Len()
}

// Take ..
func (userQ *UserQueue) Take() interface{} {
	userQ.Lock()
	v := userQ.Queue.Take()
	userQ.delCache(fmt.Sprint(v))
	userQ.Unlock()
	return v
}

// Insert ..
func (userQ *UserQueue) Insert(v interface{}) {
	vString := fmt.Sprint(v)
	if !userQ.IsExist(vString) {
		userQ.Lock()
		userQ.Queue.Insert(v)
		userQ.insertCache(vString)
		userQ.Unlock()
	}
}

// IsExist check is user already appended to user queue
func (userQ *UserQueue) IsExist(userID string) bool {
	_, found := userQ.Cache.Read(userID)
	return found
}

var userQueue *UserQueue

// DefaultUserQueue ..
func DefaultUserQueue() *UserQueue {
	return userQueue
}

func init() {
	userQueue = &UserQueue{
		Queue: module.SimpleQueue{
			List: list.New(),
		},
		Cache: module.SafeCache{
			Record: make(map[string]string),
		},
	}
}
