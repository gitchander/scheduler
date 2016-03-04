package scheduler

import (
	"fmt"
	"sync"
)

type Scheduler struct {
	open bool
	wg   *sync.WaitGroup
	quit chan struct{}
}

func Open(ts ...Task) (*Scheduler, error) {

	const formatError = "target number %d error: %s"

	for i, t := range ts {
		if t.err != nil {
			return nil, fmt.Errorf(formatError, i+1, t.err)
		}
		if t.r == nil {
			return nil, fmt.Errorf(formatError, i+1, "Runner is nil")
		}
		if t.d == nil {
			return nil, fmt.Errorf(formatError, i+1, "delayer is nil")
		}
	}

	var (
		wg   = new(sync.WaitGroup)
		quit = make(chan struct{})
	)

	wg.Add(len(ts))
	for _, t := range ts {
		go loopTask(wg, quit, t.r, t.d)
	}

	return &Scheduler{
		open: true,
		wg:   wg,
		quit: quit,
	}, nil
}

func (s *Scheduler) Close() error {

	if !s.open {
		return errorClosed
	}

	close(s.quit)
	s.wg.Wait()
	s.open = false

	return nil
}
