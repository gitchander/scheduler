package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gitchander/scheduler"
)

func main() {

	s := scheduler.New()
	defer s.Close()

	mfn := map[int]func(s *scheduler.Scheduler) error{
		0:  everySecond,
		1:  everyMinute,
		2:  everyHour,
		3:  everyDay,
		4:  everyMidnight,
		5:  everyNoon,
		6:  alarmClock,
		7:  every15SecondV1,
		8:  every15SecondV2,
		9:  secondDots,
		10: everyHalfHour,
		11: everyWeek,
	}

	fn := mfn[0]

	if err := fn(s); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Minute * 5)
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

func everySecond(s *scheduler.Scheduler) error {

	fn := func() {
		fmt.Println(time.Now())
	}

	return s.EverySecond(funcRunner{fn})
}

func everyMinute(s *scheduler.Scheduler) error {

	sec := 15

	fn := func() {
		fmt.Println(time.Now())
	}

	return s.EveryMinute(sec, funcRunner{fn})
}

func everyHour(s *scheduler.Scheduler) error {

	var (
		min = 11
		sec = 27
	)

	return s.EveryHour(min, sec, nowRunner{})
}

func everyDay(s *scheduler.Scheduler) error {

	var (
		hour = 10
		min  = 36
		sec  = 18
	)

	return s.EveryDay(hour, min, sec, nowRunner{})
}

func everyMidnight(s *scheduler.Scheduler) error {

	var (
		hour = 0
		min  = 0
		sec  = 0
	)

	return s.EveryDay(hour, min, sec, nowRunner{})
}

func everyNoon(s *scheduler.Scheduler) error {

	var (
		hour = 12
		min  = 0
		sec  = 0
	)

	return s.EveryDay(hour, min, sec, nowRunner{})
}

func alarmClock(s *scheduler.Scheduler) error {

	var (
		hour = 6
		min  = 20
		sec  = 0
	)

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
		fmt.Println("Alarm!")
	}

	return s.EveryDay(hour, min, sec, funcRunner{fn})
}

func every15SecondV1(s *scheduler.Scheduler) error {

	fn := func() {
		fmt.Println(time.Now())
	}

	for sec := 0; sec < secondsPerMinute; sec += 15 {
		if err := s.EveryMinute(sec, funcRunner{fn}); err != nil {
			return err
		}
	}
	return nil
}

func every15SecondV2(s *scheduler.Scheduler) error {

	fn := func() {
		if now := time.Now(); now.Second()%15 == 0 {
			fmt.Println(now)
		}
	}

	return s.EverySecond(funcRunner{fn})
}

func secondDots(s *scheduler.Scheduler) error {

	sec := 0

	fnSec := func() {
		if time.Now().Second() != sec {
			fmt.Print(".")
		}
	}

	fnMin := func() {
		fmt.Println(".")
	}

	if err := s.EverySecond(funcRunner{fnSec}); err != nil {
		return err
	}

	if err := s.EveryMinute(sec, funcRunner{fnMin}); err != nil {
		return err
	}

	return nil
}

func everyHalfHour(s *scheduler.Scheduler) error {

	for min := 0; min < minutesPerHour; min += 30 {
		if err := s.EveryHour(min, 0, nowRunner{}); err != nil {
			return err
		}
	}

	return nil
}

func everyWeek(s *scheduler.Scheduler) error {

	var (
		weekday = time.Wednesday
		hour    = 11
		min     = 38
		sec     = 25
	)

	fn := func() {
		now := time.Now()
		if now.Weekday() == weekday {
			fmt.Println(now)
		}
	}

	return s.EveryDay(hour, min, sec, funcRunner{fn})
}
