package main

import (
	"math/rand"

	"github.com/google/uuid"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type job struct {
	id   string
	name string
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func createJobs(amount int) []job {
	var jobs []job
	for i := 0; i < amount; i++ {
		uid := uuid.NewString()
		jobs = append(jobs, job{id: uid, name: randStringRunes(8)})
	}
	return jobs
}
