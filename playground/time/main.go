package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// build a time struct
	then := time.Date(2025, 02, 20, 18, 44, 32, 651387237, time.Local)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	// time comparison
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// get interval between two times
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// advance time forwards
	p(then.Add(diff))
	// or backwards with `-`
	p(then.Add(-diff))
}
