package runtime

import "sync"

type repository[T any] struct {
	mx   sync.RWMutex
	data map[string]T
}

func newRepository[T any]() *repository[T] {
	return &repository[T]{
		data: make(map[string]T),
	}
}

func (r *repository[T]) get(key string) (T, bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	v, ok := r.data[key]
	return v, ok
}

func (r *repository[T]) set(key string, value T) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.data[key] = value
}

func (r *repository[T]) forget(key string) bool {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, ok := r.data[key]
	if !ok {
		return false
	}
	delete(r.data, key)
	return true
}

func (r *repository[T]) clone() *repository[T] {
	var repo = newRepository[T]()
	for key, value := range r.data {
		repo.set(key, value)
	}
	return repo
}

func (r *repository[T]) foreach(fn func(key string, value T) bool) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	for key, value := range r.data {
		if !fn(key, value) {
			break
		}
	}
}
