package scheduler

import (
	"sync"
	"time"
)

type Task struct {
	err error
	r   Runner
	d   delayer
}

func loopTask(wg *sync.WaitGroup, quit <-chan struct{}, r Runner, n delayer) {

	defer wg.Done()

	for {
		d := n.getDelay()
		if d < 0 {
			return
		}

		select {
		case <-quit:
			return

		case <-time.After(d):
			go r.Run()
		}
	}
}

func EverySecond(r Runner) (t Task) {

	t.r = r
	t.d = secondDelayer{}

	return
}

func EveryMinute(sec int, r Runner) (t Task) {

	if t.err = checkSecond(sec); t.err != nil {
		return
	}

	t.r = r
	t.d = &minuteDelayer{sec}

	return
}

func EveryHour(min, sec int, r Runner) (t Task) {

	if t.err = checkMinute(min); t.err != nil {
		return
	}
	if t.err = checkSecond(sec); t.err != nil {
		return
	}

	t.r = r
	t.d = &hourDelayer{min, sec}

	return
}

func EveryDay(hour, min, sec int, r Runner) (t Task) {

	if t.err = checkHour(hour); t.err != nil {
		return
	}
	if t.err = checkMinute(min); t.err != nil {
		return
	}
	if t.err = checkSecond(sec); t.err != nil {
		return
	}

	t.r = r
	t.d = &dayDelayer{hour, min, sec}

	return
}
