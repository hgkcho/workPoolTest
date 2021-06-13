package main

import (
	"testing"

	"github.com/google/uuid"
)

func Benchmark_Pool(t *testing.B) {
	d := newDispatchar()
	d.Start()
	for _, v := range createJobs(JobCount) {
		d.Add(v)
	}
	d.Wait()
}
func Benchmark_Normarl(t *testing.B) {
	w := worker{
		id: uuid.NewString(),
	}
	for _, v := range createJobs(JobCount) {
		w.DoWork(v)
	}
}
