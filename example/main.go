package main

import (
	"fmt"
	"time"

	"github.com/gitchander/scheduler"
)

func main() {
	everySecond()
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

func everySecond() {

	fn := func() {
		fmt.Println(time.Now())
	}

	t := scheduler.Task{
		Runner: funcRunner{fn},
		Nexter: scheduler.SecondNexter{},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(20 * time.Second)
}

func everyMinute() {

	fn := func() {
		fmt.Println(time.Now())
	}

	t := scheduler.Task{
		Runner: funcRunner{fn},
		Nexter: scheduler.MinuteNexter{
			Second: 15,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(10 * time.Minute)
}

func everyHour() {

	t := scheduler.Task{
		Runner: nowRunner{},
		Nexter: scheduler.HourNexter{
			Minute: 11,
			Second: 27,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(5 * time.Minute)
}

func everyDay() {

	t := scheduler.Task{
		Runner: nowRunner{},
		Nexter: scheduler.DayNexter{
			Hour:   17,
			Minute: 43,
			Second: 21,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(5 * time.Minute)
}

func everyMidnight() {

	t := scheduler.Task{
		Runner: nowRunner{},
		Nexter: scheduler.DayNexter{
			Hour:   0,
			Minute: 0,
			Second: 0,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(5 * time.Minute)
}

func everyNoon() {

	t := scheduler.Task{
		Runner: nowRunner{},
		Nexter: scheduler.DayNexter{
			Hour:   12,
			Minute: 0,
			Second: 0,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(5 * time.Minute)
}

func every15SecondV1() {

	fn := func() {
		fmt.Println(time.Now())
	}

	var ts []scheduler.Task

	for sec := 0; sec < secondsPerMinute; sec += 15 {

		t := scheduler.Task{
			Runner: funcRunner{fn},
			Nexter: scheduler.MinuteNexter{
				Second: sec,
			},
		}

		ts = append(ts, t)
	}

	s := scheduler.Open(ts...)
	defer s.Close()

	time.Sleep(3 * time.Minute)
}

func every15SecondV2() {

	fn := func() {
		if now := time.Now(); now.Second()%15 == 0 {
			fmt.Println(now)
		}
	}

	t := scheduler.Task{
		Runner: funcRunner{fn},
		Nexter: scheduler.SecondNexter{},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(3 * time.Minute)
}

func secondDots() {

	sec := 0

	fnSec := func() {
		if time.Now().Second() != sec {
			fmt.Print(".")
		}
	}

	fnMin := func() {
		fmt.Println(".")
	}

	taskSec := scheduler.Task{
		Runner: funcRunner{fnSec},
		Nexter: scheduler.SecondNexter{},
	}

	taskMin := scheduler.Task{
		Runner: funcRunner{fnMin},
		Nexter: scheduler.MinuteNexter{
			Second: sec,
		},
	}

	s := scheduler.Open(taskSec, taskMin)
	defer s.Close()

	time.Sleep(10 * time.Minute)
}

func alarmClock() {

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

	t := scheduler.Task{
		Runner: funcRunner{fn},
		Nexter: scheduler.DayNexter{
			Hour:   18,
			Minute: 21,
			Second: 0,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(10 * time.Minute)
}

func everyHalfHour() {

	var ts []scheduler.Task

	for min := 0; min < minutesPerHour; min += 30 {

		t := scheduler.Task{
			Runner: nowRunner{},
			Nexter: scheduler.HourNexter{
				Minute: min,
				Second: 0,
			},
		}

		ts = append(ts, t)
	}

	s := scheduler.Open(ts...)
	defer s.Close()

	time.Sleep(3 * time.Hour)
}

func everyWeek() {

	var weekday = time.Wednesday

	fn := func() {
		now := time.Now()
		if now.Weekday() == weekday {
			fmt.Println(now)
		}
	}

	t := scheduler.Task{
		Runner: funcRunner{fn},
		Nexter: scheduler.DayNexter{
			Hour:   11,
			Minute: 38,
			Second: 25,
		},
	}

	s := scheduler.Open(t)
	defer s.Close()

	time.Sleep(5 * time.Minute)
}
