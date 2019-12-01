package balancer

import (
	"testing"
)

func TestRoundRobin_Pick(t *testing.T) {

	rb := New([]interface{}{1, 2, 3, 4, 5})

	for i := 0; i < 1000; i++ {
		if item, err := rb.Pick(); err != nil {
			t.FailNow()
		} else {
			t.Log(item.(int))
		}
	}
}
