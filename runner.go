package scheduler

import (
	"sync"
	"time"
)

type Runner interface {
	Run()
}

func loopTask(wg *sync.WaitGroup, quit <-chan struct{}, n nexter, r Runner) {

	defer wg.Done()

	for {
		var (
			now  = time.Now()
			next = n.Next(now)
			d    = next.Sub(now)
		)

		select {
		case <-quit:
			return

		case <-time.After(d):
			go r.Run()
		}
	}
}
