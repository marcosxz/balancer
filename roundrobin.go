package balancer

import (
	"errors"
	"sync"
)

var ErrNoAvailableItem = errors.New("roundRobin no item is available")

// Balance roundRobin instance
type RoundRobin struct {
	m     sync.Mutex
	next  int
	items []interface{}
}

// New balance instance
func New(items []interface{}) *RoundRobin {
	return &RoundRobin{items: items}
}

// Pick available item
func (b *RoundRobin) Pick() (interface{}, error) {
	if len(b.items) == 0 {
		return nil, ErrNoAvailableItem
	}

	b.m.Lock()
	r := b.items[b.next]
	b.next = (b.next + 1) % len(b.items)
	b.m.Unlock()

	return r, nil
}
