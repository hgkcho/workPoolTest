package main

import (
	"strconv"
	"sync"
)

const MaxWorkerPool = 10
const MaxJobQueue = 1000

type dispatchar struct {
	jobQueue   chan job
	workerPool chan *worker
	workers    []*worker
	wg         sync.WaitGroup
	end        chan struct{}
}

func newDispatchar() *dispatchar {
	return &dispatchar{
		jobQueue:   make(chan job, MaxJobQueue),
		workerPool: make(chan *worker, MaxWorkerPool),
		workers:    make([]*worker, MaxWorkerPool),
		wg:         sync.WaitGroup{},
		end:        make(chan struct{}),
	}
}

func (d *dispatchar) Start() {
	for i := 0; i < cap(d.workerPool); i++ {
		w := worker{
			id:         strconv.Itoa(i),
			jobQueue:   make(chan job),
			dispatchar: d,
			end:        make(chan struct{}),
		}
		w.Start()
		d.workers = append(d.workers, &w)
	}

	go func() {
		for {
			select {
			case <-d.end:
				for _, v := range d.workers {
					v.Stop()
				}
			case job := <-d.jobQueue:
				w := <-d.workerPool
				w.jobQueue <- job
			}
		}
	}()
}

func (d *dispatchar) Add(j job) {
	d.wg.Add(1)
	d.jobQueue <- j
}

func (d *dispatchar) Wait() {
	d.wg.Wait()
}
