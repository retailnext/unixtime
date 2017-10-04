// Copyright (c) 2017, RetailNext, Inc.

package unixtime

import (
	"math"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	cases := []struct {
		unix     int64
		duration time.Duration
		time     time.Time
	}{
		{0, time.Second, time.Unix(0, 0)},
		{1472515778, time.Second, time.Unix(1472515778, 0)},
		{1472515778000, time.Millisecond, time.Unix(1472515778, 0)},
		{1472515778000000, time.Microsecond, time.Unix(1472515778, 0)},
		{1472515778000000000, time.Nanosecond, time.Unix(1472515778, 0)},

		// this will overflow if you just use the time nanosecond functions
		{147251577800, time.Second, time.Unix(147251577800, 0)},
		{14725157780000, time.Second, time.Unix(14725157780000, 0)},
		{1472515778000000, time.Second, time.Unix(1472515778000000, 0)},
		{147251577800000000, time.Second, time.Unix(147251577800000000, 0)},

		{math.MaxInt64, time.Millisecond, time.Unix(9223372036854775, 807000000)},
		{math.MaxInt64, time.Microsecond, time.Unix(9223372036854, 775807000)},
	}

	for _, tc := range cases {
		t0 := ToTime(tc.unix, tc.duration)
		if !t0.Equal(tc.time) {
			t.Errorf("%d %s %s != %s", tc.unix, tc.duration, tc.time, t0)
		}
		gotUnix := ToUnix(t0, tc.duration)
		if gotUnix != tc.unix {
			t.Errorf("%d %s != %d", tc.unix, tc.duration, gotUnix)
		}
	}
}
