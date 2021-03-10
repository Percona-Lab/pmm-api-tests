package pmmapitests

import (
	"math/rand"
	"sync"
)

type ConcurrentRand struct {
	m    sync.Mutex
	rand *rand.Rand
}

func NewConcurrentRand(seed int64) *ConcurrentRand {
	r := &ConcurrentRand{
		rand: rand.New(rand.NewSource(seed)),
	}
	return r
}
func (r *ConcurrentRand) Int63() int64 {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rand.Int63()
}

func (r *ConcurrentRand) Seed(seed int64) {
	r.m.Lock()
	defer r.m.Unlock()
	r.rand.Seed(seed)
}

func (r *ConcurrentRand) Uint64() uint64 {
	r.m.Lock()
	defer r.m.Unlock()
	return r.rand.Uint64()
}
