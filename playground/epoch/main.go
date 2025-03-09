// get time since unix epoch

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// use `time.Now` with `Unix`, `UnixMilli` or `UnixNano` to get elapsed
	// time since unix epoch in seconds, milliseconds and nanoseconds respectively
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// convert integer seconds or nanoseconds since the epoch into the
	// corresponding time
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
