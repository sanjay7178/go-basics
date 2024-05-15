package main

import (
	// "fmt"
	// "sync"
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32

	//  go x
	// go y
	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state , int32(i))
		}(i)
	}
}
