package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"time"
)

type worker struct {
	id         string
	jobQueue   chan job
	dispatchar *dispatchar
	end        chan struct{}
}

func (w worker) DoWork(j job) {
	h := fnv.New32a()
	h.Write([]byte(j.name))
	time.Sleep(time.Second)
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("worker [%s] - created hash [%d] from word [%s]\n", w.id, h.Sum32(), j.name)
	}
}

func (w *worker) Start() {
	go func() {
		for {
			w.dispatchar.workerPool <- w
			select {
			case job := <-w.jobQueue:
				w.DoWork(job)
				w.dispatchar.wg.Done()
			case <-w.end:
				return
			}
		}
	}()
}

func (w *worker) Stop() {
	w.end <- struct{}{}
}
