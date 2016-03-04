package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gitchander/scheduler"
)

func main() {
	if err := everySecond(); err != nil {
		log.Fatal(err)
	}
}

const (
	daysPerWeek = 7

	hoursPerDay  = 24
	hoursPerWeek = hoursPerDay * daysPerWeek

	minutesPerHour = 60
	minutesPerDay  = minutesPerHour * hoursPerDay
	minutesPerWeek = minutesPerDay * daysPerWeek

	secondsPerMinute = 60
	secondsPerHour   = secondsPerMinute * minutesPerHour
	secondsPerDay    = secondsPerHour * hoursPerDay
	secondsPerWeek   = secondsPerDay * daysPerWeek
)

type funcRunner struct {
	fn func()
}

func (r funcRunner) Run() {
	r.fn()
}

type nowRunner struct{}

func (nowRunner) Run() {
	fmt.Println(time.Now())
}

func everySecond() error {

	fn := func() {
		fmt.Println(time.Now())
	}

	s, err := scheduler.Open(
		scheduler.EverySecond(funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(20 * time.Second)

	return nil
}

func everyMinute() error {

	sec := 27

	fn := func() {
		fmt.Println(time.Now())
	}

	s, err := scheduler.Open(
		scheduler.EveryMinute(sec, funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(10 * time.Minute)

	return nil
}

func everyHour() error {

	var (
		min = 11
		sec = 27
	)

	s, err := scheduler.Open(
		scheduler.EveryHour(min, sec, nowRunner{}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(5 * time.Minute)

	return nil
}

func everyDay() error {

	var (
		hour = 17
		min  = 43
		sec  = 21
	)

	s, err := scheduler.Open(
		scheduler.EveryDay(hour, min, sec, nowRunner{}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(5 * time.Minute)

	return nil
}

func everyMidnight() error {

	var (
		hour = 0
		min  = 0
		sec  = 0
	)

	t := scheduler.EveryDay(hour, min, sec, nowRunner{})

	s, err := scheduler.Open(t)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(5 * time.Minute)

	return nil
}

func everyNoon() error {

	var (
		hour = 12
		min  = 0
		sec  = 0
	)

	t := scheduler.EveryDay(hour, min, sec, nowRunner{})

	s, err := scheduler.Open(t)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(5 * time.Minute)

	return nil
}

func every15SecondV1() error {

	fn := func() {
		fmt.Println(time.Now())
	}

	s, err := scheduler.Open(
		scheduler.EveryMinute(0, funcRunner{fn}),
		scheduler.EveryMinute(15, funcRunner{fn}),
		scheduler.EveryMinute(30, funcRunner{fn}),
		scheduler.EveryMinute(45, funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(3 * time.Minute)

	return nil
}

func every15SecondV2() error {

	fn := func() {
		if now := time.Now(); now.Second()%15 == 0 {
			fmt.Println(now)
		}
	}

	s, err := scheduler.Open(
		scheduler.EverySecond(funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(3 * time.Minute)

	return nil
}

func secondDots() error {

	sec := 0

	fnSec := func() {
		if time.Now().Second() != sec {
			fmt.Print(".")
		}
	}

	fnMin := func() {
		fmt.Println(".")
	}

	s, err := scheduler.Open(
		scheduler.EverySecond(funcRunner{fnSec}),
		scheduler.EveryMinute(sec, funcRunner{fnMin}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(10 * time.Minute)

	return nil
}

func alarmClock() error {

	fn := func() {
		now := time.Now()
		switch weekDay := now.Weekday(); weekDay {
		case time.Monday:
		case time.Tuesday:
		case time.Wednesday:
		case time.Thursday:
		case time.Friday:
		default:
			return
		}
		fmt.Println(time.Now(), "Alarm!")
	}

	var (
		hour = 12
		min  = 47
		sec  = 17
	)

	s, err := scheduler.Open(
		scheduler.EveryDay(hour, min, sec, funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(10 * time.Minute)

	return nil
}

func everyHalfHour() error {

	var ts []scheduler.Task

	const sec = 0

	for min := 0; min < minutesPerHour; min += 30 {
		t := scheduler.EveryHour(min, sec, nowRunner{})
		ts = append(ts, t)
	}

	s, err := scheduler.Open(ts...)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(3 * time.Hour)

	return nil
}

func everyWeek() error {

	var weekday = time.Wednesday

	fn := func() {
		if now := time.Now(); now.Weekday() == weekday {
			fmt.Println(now)
		}
	}

	var (
		hour = 18
		min  = 21
		sec  = 17
	)

	s, err := scheduler.Open(
		scheduler.EveryDay(hour, min, sec, funcRunner{fn}),
	)
	if err != nil {
		return err
	}
	defer s.Close()

	time.Sleep(5 * time.Minute)

	return nil
}
