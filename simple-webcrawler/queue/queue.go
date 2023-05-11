package queue

import "sync"

type Queue struct {
	maxThreads int
	jobs       chan func()
	wg         sync.WaitGroup
}

func NewQueue(maxThreads int) *Queue {
	return &Queue{
		maxThreads: maxThreads,
		jobs:       make(chan func()),
	}
}

func (q *Queue) Start() {
	for i := 0; i < q.maxThreads; i++ {
		go func() {
			for job := range q.jobs {
				job()
				q.wg.Done()
			}
		}()
	}
}

func (q *Queue) AddJob(job func()) {
	q.wg.Add(1)
	go func() {
		q.jobs <- job
	}()
}

func (q *Queue) Wait() {
	q.wg.Wait()
}
