package scheduler

import "sync"

type Scheduler struct {
	opened bool
	wg     *sync.WaitGroup
	quit   chan struct{}
}

func New() *Scheduler {

	return &Scheduler{
		opened: true,
		wg:     new(sync.WaitGroup),
		quit:   make(chan struct{}),
	}
}

func (s *Scheduler) Close() error {

	if !s.opened {
		return errorClosed
	}

	close(s.quit)
	s.wg.Wait()
	s.opened = false

	return nil
}

func (s *Scheduler) EverySecond(r Runner) error {

	if !s.opened {
		return errorClosed
	}

	n := secondNexter{}
	return s.every(n, r)
}

func (s *Scheduler) EveryMinute(sec int, r Runner) error {

	if !s.opened {
		return errorClosed
	}

	if err := checkSecond(sec); err != nil {
		return err
	}

	n := &minuteNexter{sec}
	return s.every(n, r)
}

func (s *Scheduler) EveryHour(min, sec int, r Runner) error {

	if !s.opened {
		return errorClosed
	}

	if err := checkMinute(min); err != nil {
		return err
	}
	if err := checkSecond(sec); err != nil {
		return err
	}

	n := &hourNexter{min, sec}
	return s.every(n, r)
}

func (s *Scheduler) EveryDay(hour, min, sec int, r Runner) error {

	if !s.opened {
		return errorClosed
	}

	if err := checkHour(hour); err != nil {
		return err
	}
	if err := checkMinute(min); err != nil {
		return err
	}
	if err := checkSecond(sec); err != nil {
		return err
	}

	n := &dayNexter{hour, min, sec}
	return s.every(n, r)
}

func (s *Scheduler) every(n nexter, r Runner) error {

	s.wg.Add(1)
	go loopTask(s.wg, s.quit, n, r)

	return nil
}
