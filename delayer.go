package scheduler

import "time"

type delayer interface {
	getDelay() time.Duration
}

type secondDelayer struct{}

func (secondDelayer) getDelay() time.Duration {

	now := time.Now()

	next := now.Truncate(time.Second)

	if next.Before(now) {
		next = next.Add(time.Second)
	}

	return next.Sub(now)
}

type minuteDelayer struct {
	sec int
}

func (n *minuteDelayer) getDelay() time.Duration {

	now := time.Now()

	next := now.Truncate(time.Minute)
	next = next.Add(time.Second * time.Duration(n.sec))

	if next.Before(now) {
		next = next.Add(time.Minute)
	}

	return next.Sub(now)
}

type hourDelayer struct {
	min, sec int
}

func (n *hourDelayer) getDelay() time.Duration {

	now := time.Now()

	next := now.Truncate(time.Hour)
	next = next.Add(time.Minute*time.Duration(n.min) +
		time.Second*time.Duration(n.sec))

	if next.Before(now) {
		next = next.Add(time.Hour)
	}

	return next.Sub(now)
}

type dayDelayer struct {
	hour, min, sec int
}

func (n *dayDelayer) getDelay() time.Duration {

	now := time.Now()

	year, month, day := now.Date()
	next := time.Date(year, month, day, n.hour, n.min, n.sec, 0, now.Location())

	if next.Before(now) {
		next = next.Add(time.Hour * 24)
	}

	return next.Sub(now)
}
