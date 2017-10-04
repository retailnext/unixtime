# unixtime

unixtime is a go package that makes it easier to convert between unix timestamps and go objects. Its primary purpose is so you don't have to remember (and properly type) the number of nanoseconds in a millisecond. It also properly handles the seconds/nanoseconds division in time objects so it avoids overflows of times above math.MaxInt64 nanoseconds since 1 January 1970 (times after year 2262).

### Documentation

http://godoc.org/github.com/retailnext/unixtime

### Usage

    t0 := time.Now()
    unixMicroseconds := unixtime.ToUnix(t0, time.Microsecond)

    fmt.Printf("Microseconds: %d\n", unixMicroseconds)
    // prints Microseconds: 1507158285672039

    t1 := unixtime.ToTime(unixMicroseconds, time.Microsecond)

    fmt.Printf("t0=%q t1=%q\n", t0.Format(time.RFC3339Nano), t1.Format(time.RFC3339Nano))
    // prints t0="2017-10-04T23:04:45.672039138Z" t1="2017-10-04T23:04:45.672039Z"


### Installation

    go get github.com/retailnext/unixtime
