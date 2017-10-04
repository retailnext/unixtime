// Copyright (c) 2017, RetailNext, Inc.

package unixtime

import (
	"fmt"
	"time"
)

func ExampleToUnix() {
	now := time.Now()

	s := ToUnix(now, time.Second)
	ms := ToUnix(now, time.Millisecond)
	us := ToUnix(now, time.Microsecond)

	fmt.Printf("seconds=%d milliseconds=%d microseconds=%d\n", s, ms, us)
}
