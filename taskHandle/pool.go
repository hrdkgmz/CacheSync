package taskHandle

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type TaskHandler func() error

type WorkPool struct {
	closed       int32
	isQueTask    int32
	errChan      chan error
	timeout      time.Duration
	wg           sync.WaitGroup
	task         chan TaskHandler
	waitingQueue *TaskQueue
}

func NewPool(max int, timeOut time.Duration) *WorkPool {
	if max < 1 {
		max = 1
	}

	p := &WorkPool{
		task:         make(chan TaskHandler, 2*max),
		errChan:      make(chan error, 1),
		waitingQueue: NewQueue(),
		timeout:      timeOut,
	}

	go p.loop(max)
	return p
}

func (p *WorkPool) SetTimeout(timeout time.Duration) {
	p.timeout = timeout
}

func (p *WorkPool) Do(fn TaskHandler) {
	if p.IsClosed() {
		return
	}
	p.waitingQueue.Push(fn)
}

func (p *WorkPool) DoWait(task TaskHandler) {
	if p.IsClosed() {
		return
	}

	doneChan := make(chan struct{})
	p.waitingQueue.Push(TaskHandler(func() error {
		defer close(doneChan)
		return task()
	}))
	<-doneChan
}

func (p *WorkPool) Wait() error {
	p.waitingQueue.Wait()
	p.waitingQueue.Close()
	p.waitTask()
	close(p.task)
	p.wg.Wait()
	select {
	case err := <-p.errChan:
		return err
	default:
		return nil
	}
}

func (p *WorkPool) IsDone() bool {
	if p == nil || p.task == nil {
		return true
	}

	return p.waitingQueue.Len() == 0 && len(p.task) == 0
}

func (p *WorkPool) IsClosed() bool {
	if atomic.LoadInt32(&p.closed) == 1 {
		return true
	}
	return false
}

func (p *WorkPool) startQueue() {
	p.isQueTask = 1
	for {
		tmp := p.waitingQueue.Pop()
		if p.IsClosed() {
			p.waitingQueue.Close()
			break
		}
		if tmp != nil {
			fn := tmp.(TaskHandler)
			if fn != nil {
				p.task <- fn
			}
		} else {
			break
		}

	}
	atomic.StoreInt32(&p.isQueTask, 0)
}

func (p *WorkPool) waitTask() {
	for {
		runtime.Gosched()
		if p.IsDone() {
			if atomic.LoadInt32(&p.isQueTask) == 0 {
				break
			}
		}
	}
}

func (p *WorkPool) loop(maxWorkersCount int) {
	go p.startQueue()

	p.wg.Add(maxWorkersCount)

	for i := 0; i < maxWorkersCount; i++ {
		go func() {
			defer p.wg.Done()

			for wt := range p.task {
				if wt == nil || atomic.LoadInt32(&p.closed) == 1 {
					continue
				}

				closed := make(chan struct{}, 1)

				if p.timeout > 0 {
					ct, cancel := context.WithTimeout(context.Background(), p.timeout)
					go func() {
						select {
						case <-ct.Done():
							p.errChan <- ct.Err()
							atomic.StoreInt32(&p.closed, 1)
							cancel()
						case <-closed:
						}
					}()
				}

				err := wt()
				close(closed)
				if err != nil {
					select {
					case p.errChan <- err:
						atomic.StoreInt32(&p.closed, 1)
					default:
					}
				}
			}
		}()
	}
}
