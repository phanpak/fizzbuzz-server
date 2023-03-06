package fizzbuzz

import (
	"sync"
)

type Counter struct {
	hits map[string]*Count
}

type Count struct {
	count int
	mutex sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{hits: make(map[string]*Count)}
}

func (h *Counter) Increment(key string) {
	hit, ok := h.hits[key]
	if !ok {
		hit = &Count{}
		h.hits[key] = hit
	}
	hit.mutex.Lock()
	hit.count++
	hit.mutex.Unlock()
}

type MaxResult struct {
	Key   string
	Count int
}

func (h *Counter) Max() MaxResult {
	max := MaxResult{"{}", 0}
	for key, hit := range h.hits {
		hit.mutex.Lock()
		if hit.count > max.Count {
			max = MaxResult{key, hit.count}
		}
		hit.mutex.Unlock()
	}
	return max
}
