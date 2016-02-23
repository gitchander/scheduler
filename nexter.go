package scheduler

import "time"

type nexter interface {
	Next(t time.Time) time.Time
}

type secondNexter struct{}

func (secondNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Second)

	if next.Before(t) {
		next = next.Add(time.Second)
	}

	return next
}

type minuteNexter struct {
	sec int
}

func (n *minuteNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Minute)
	next = next.Add(time.Second * time.Duration(n.sec))

	if next.Before(t) {
		next = next.Add(time.Minute)
	}

	return next
}

type hourNexter struct {
	min, sec int
}

func (n *hourNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Hour)
	next = next.Add(time.Minute*time.Duration(n.min) +
		time.Second*time.Duration(n.sec))

	if next.Before(t) {
		next = next.Add(time.Hour)
	}

	return next
}

type dayNexter struct {
	hour, min, sec int
}

func (n *dayNexter) Next(t time.Time) time.Time {

	year, month, day := t.Date()
	next := time.Date(year, month, day, n.hour, n.min, n.sec, 0, t.Location())

	if next.Before(t) {
		next = next.Add(time.Hour * 24)
	}

	return next
}
