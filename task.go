package scheduler

import (
	"sync"
	"time"
)

type Task struct {
	Runner
	Nexter
}

func loopTask(wg *sync.WaitGroup, quit <-chan struct{}, n Nexter, r Runner) {

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

/*
func EverySecond(r Runner) *Task {
	return &Task{r, SecondNexter{}}
}

func EveryMinute(sec int, r Runner) (*Task, error) {

	if err := checkSecond(sec); err != nil {
		return nil, err
	}

	return &Task{r, &MinuteNexter{sec}}, nil
}

func EveryHour(min, sec int, r Runner) (*Task, error) {

	if err := checkMinute(min); err != nil {
		return nil, err
	}
	if err := checkSecond(sec); err != nil {
		return nil, err
	}

	return &Task{r, &HourNexter{min, sec}}, nil
}

func EveryDay(hour, min, sec int, r Runner) (*Task, error) {

	if err := checkHour(hour); err != nil {
		return nil, err
	}
	if err := checkMinute(min); err != nil {
		return nil, err
	}
	if err := checkSecond(sec); err != nil {
		return nil, err
	}

	return &Task{r, &DayNexter{hour, min, sec}}, nil
}
*/
