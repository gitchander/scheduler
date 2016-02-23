package scheduler

import "errors"

var errorClosed = errors.New("scheduler is not created or closed")

func checkHour(hour int) error {
	if (hour < 0) || (hour > 23) {
		return errors.New("hour must be [0 ... 23]")
	}
	return nil
}

func checkMinute(min int) error {
	if (min < 0) || (min > 59) {
		return errors.New("minute must be [0 ... 59]")
	}
	return nil
}

func checkSecond(sec int) error {
	if (sec < 0) || (sec > 59) {
		return errors.New("second must be [0 ... 59]")
	}
	return nil
}
