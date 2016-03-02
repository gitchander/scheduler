package scheduler

import "time"

type Nexter interface {
	Next(t time.Time) time.Time
}

type SecondNexter struct{}

func (SecondNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Second)

	if next.Before(t) {
		next = next.Add(time.Second)
	}

	return next
}

type MinuteNexter struct {
	Second int
}

func (n MinuteNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Minute)
	next = next.Add(time.Second * time.Duration(n.Second))

	if next.Before(t) {
		next = next.Add(time.Minute)
	}

	return next
}

type HourNexter struct {
	Minute, Second int
}

func (n HourNexter) Next(t time.Time) time.Time {

	next := t.Truncate(time.Hour)
	next = next.Add(time.Minute*time.Duration(n.Minute) +
		time.Second*time.Duration(n.Second))

	if next.Before(t) {
		next = next.Add(time.Hour)
	}

	return next
}

type DayNexter struct {
	Hour, Minute, Second int
}

func (n DayNexter) Next(t time.Time) time.Time {

	year, month, day := t.Date()
	next := time.Date(year, month, day, n.Hour, n.Minute, n.Second, 0, t.Location())

	if next.Before(t) {
		next = next.Add(time.Hour * 24)
	}

	return next
}
