package scheduler

import "sync"

type Scheduler struct {
	open bool
	wg   *sync.WaitGroup
	quit chan struct{}
}

func Open(ts ...Task) *Scheduler {

	s := &Scheduler{
		open: true,
		wg:   new(sync.WaitGroup),
		quit: make(chan struct{}),
	}

	s.wg.Add(len(ts))

	for _, t := range ts {
		go loopTask(s.wg, s.quit, t.Nexter, t.Runner)
	}

	return s
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
