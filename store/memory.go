package store

import "sync"

type Memory[T any] struct {
	mtx    sync.RWMutex
	nextID int
	data   []T
}

func NewMemory[T any]() *Memory[T] {
	return &Memory[T]{data: make([]T, 0), nextID: 1}
}

func (m *Memory[T]) WithID(setID func(*T, int)) func(T) T {
	return func(t T) T {
		m.mtx.Lock()
		defer m.mtx.Unlock()
		setID(&t, m.nextID)
		m.nextID++
		m.data = append(m.data, t)
		return t
	}
}

func (m *Memory[T]) All() []T {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]T, len(m.data))
	copy(out, m.data)
	return out
}

func (m *Memory[T]) Delete(match func(T) bool) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for i, v := range m.data {
		if match(v) {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return true
		}
	}
	return false
}

func (m *Memory[T]) Update(match func(*T) bool) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for i := range m.data {
		if match(&m.data[i]) {
			return true
		}
	}
	return false
}
