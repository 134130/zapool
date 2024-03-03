package zapool

import "sync"

// New creates a new pool of objects.
func New[T any](new func() T) Pool[T] {
	return &pool[T]{
		p: &sync.Pool{
			New: func() interface{} {
				return new()
			},
		},
	}
}

// Pool is a generic pool of objects.
type Pool[T any] interface {
	// Get returns an object from the pool.
	Get() T
	// Put returns an object to the pool.
	Put(T)
}

type pool[T any] struct {
	p *sync.Pool
}

func (p *pool[T]) Get() T {
	return p.p.Get().(T)
}

func (p *pool[T]) Put(obj T) {
	p.p.Put(obj)
}
