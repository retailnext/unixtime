// Copyright (c) 2017, RetailNext, Inc.

package unixtime

import (
	"fmt"
	"time"
)

func ExampleToTime() {
	var unixMilliseconds int64 = 1507156635642

	t := ToTime(unixMilliseconds, time.Second)

	fmt.Printf("time=%s\n", t.Format(time.RFC3339Nano))
}
