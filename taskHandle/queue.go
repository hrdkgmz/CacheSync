package taskHandle

import (
	"runtime"
	"sync"
	"sync/atomic"

	"gopkg.in/eapache/queue.v1"
)

type TaskQueue struct {
	sync.Mutex
	popable *sync.Cond
	buffer  *queue.Queue
	closed  bool
	count   int32
}

func NewQueue() *TaskQueue {

	ch := &TaskQueue{
		buffer: queue.New(),
	}
	ch.popable = sync.NewCond(&ch.Mutex)
	return ch
}

func (q *TaskQueue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for q.Len() == 0 && !q.closed {
		c.Wait()
	}

	if q.closed { //已关闭
		return
	}

	if q.Len() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		atomic.AddInt32(&q.count, -1)
	}
	return
}

func (q *TaskQueue) TryPop() (v interface{}, ok bool) {
	buffer := q.buffer

	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if q.Len() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		atomic.AddInt32(&q.count, -1)
		ok = true
	} else if q.closed {
		ok = true
	}

	return
}

func (q *TaskQueue) Push(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		q.buffer.Add(v)
		atomic.AddInt32(&q.count, 1)
		q.popable.Signal()
	}
}

func (q *TaskQueue) Len() int {
	return (int)(atomic.LoadInt32(&q.count))
}

func (q *TaskQueue) Close() {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if !q.closed {
		q.closed = true
		atomic.StoreInt32(&q.count, 0)
		q.popable.Broadcast() //广播
	}
}

func (q *TaskQueue) Wait() {
	for {
		if q.closed || q.Len() == 0 {
			break
		}

		runtime.Gosched()
	}
}
