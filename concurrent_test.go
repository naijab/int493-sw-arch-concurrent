package concurrent

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCountWithMutex(t *testing.T) {
	// if have mutex more than one it free for each
	l := sync.Mutex{}
	count := 0

	for i := 0; i < 1000; i++ {
		go func() {

			for k := 0; k < 1000; k++ {
				l.Lock()
				count++
				l.Unlock()
			}
		}()
	}

	time.Sleep(1 * time.Second)

	if count != 1000000 {
		t.Errorf("actual=%d", count)
	}
}

func TestCountWithAtomic(t *testing.T) {
	var count int32 = 0

	for i := 0; i < 1000; i++ {
		go func() {

			for k := 0; k < 1000; k++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}

	time.Sleep(1 * time.Second)

	if count != 1000000 {
		t.Errorf("actual=%d", count)
	}
}

func TestCountWithWaitingGroup(t *testing.T) {
	var count int32 = 0

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for k := 0; k < 1000; k++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait() // lock and wait util 1000 goroutine finish

	if count != 1000000 {
		t.Errorf("actual=%d", count)
	}
}