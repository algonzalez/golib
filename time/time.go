package time

import "time"

// SleepSeconds pauses the current goroutine for at least n seconds.
// A negative or zero duration causes SleepSeconds to return immediately.
func SleepSeconds(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

// SleepMicroseconds pauses the current goroutine for at least n milliseconds.
// A negative or zero duration causes SleepMicroseconds to return immediately.
func SleepMicroseconds(n int) {
	time.Sleep(time.Duration(n) * time.Microsecond)
}

// SleepMilliseconds pauses the current goroutine for at least n milliseconds.
// A negative or zero duration causes SleepMilliseconds to return immediately.
func SleepMilliseconds(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// TimeIt returns the time elapsed while running action.
func TimeIt(action func()) (elapsed time.Duration) {
	start := time.Now()
	action()
	return time.Now().Sub(start)
}
