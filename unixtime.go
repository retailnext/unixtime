// Copyright (c) 2017, RetailNext, Inc.

// Package unixtime provides functions that make it easier to convert between go time.Time objects and Unix timestamps.
package unixtime

import "time"

// ToUnix converts a time object to number of
// {Nano,Mirco,Milli,}seconds since 1 January 1970 utc.
// ToUnix panics if your unix is something other than
// 1 {Nano,Micro,Milli}second.
func ToUnix(t time.Time, unit time.Duration) int64 {
	var secMult int64
	var nanoDiv int64

	switch unit {
	case time.Nanosecond:
		secMult = 1000000000
		nanoDiv = 1
	case time.Microsecond:
		secMult = 1000000
		nanoDiv = 1000
	case time.Millisecond:
		secMult = 1000
		nanoDiv = 1000000
	case time.Second:
		secMult = 1
		nanoDiv = 1000000000
	default:
		panic("unit must be one of Second,Millisecond,Microsecond,Nanosecond")
	}

	unix := t.Unix() * secMult
	nano := int64(t.Nanosecond()) / nanoDiv

	return unix + nano
}

// ToTime converts an int64 representing the number
// of unit {nano,mirco,milli,}seconds since 1 January 1970 utc.
// to a time. This function has the same imput limitations as
// time.Unix.
// ToTime panics if your unix is something other than
// 1 {Nano,Micro,Milli}second.
func ToTime(unixTime int64, unit time.Duration) time.Time {
	var sec int64
	var nsec int64

	switch unit {
	case time.Nanosecond:
		sec = unixTime / 1000000000
		nsec = unixTime % 1000000000
	case time.Microsecond:
		sec = unixTime / 1000000
		nsec = (unixTime % 1000000) * 1000
	case time.Millisecond:
		sec = unixTime / 1000
		nsec = (unixTime % 1000) * 1000000
	case time.Second:
		sec = unixTime
	default:
		panic("unit must be one of Second,Millisecond,Microsecond,Nanosecond")
	}

	return time.Unix(sec, nsec)
}
