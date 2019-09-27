package goutils

import (
	"time"
)

var (
	MinTime = time.Date(1971,0,0,0,0,0,0,time.UTC)
	MaxTime = time.Date(2100,0,0,0,0,0,0,time.UTC)
)

// UnixMilli creates time.Time from milliseconds time stamp ie 1523098153599
func UnixMilli(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}

// Time alias for `time.Time`
type Time time.Time

// UnixMilli number of miliseconds since 1970
func (t Time) UnixMilli() int64 {
	return (time.Time)(t).UnixNano() / int64(time.Millisecond)
}

// TimeOrMinTime returns `MinTime` if `t` is nil
func TimeOrMinTime(t *time.Time) *time.Time {
	if t == nil {
		return &MinTime
	}
	return t
}

// TimeOrNow returns `time.Now` if `t` is nil or `t` is greater than `time.Now`
func TimeClipToNow(t *time.Time) *time.Time {
	if t == nil || t.After(time.Now()) {
		now := time.Now()
		return  &now
	}
	return t
}

// NOTE: I think below is no longer used. Delete and make sure cyrus compiles

// UnixFloatNano from `Time`
func (t Time) UnixFloatNano() float64 {
	return float64(time.Time(t).UnixNano()) / float64(time.Second)
}

type Duration time.Duration

func (d Duration) UnixFloatNano() float64 {
	return float64(time.Duration(d).Nanoseconds()) / float64(time.Second)
}

func (t Time)slqTimeStr() string {
	return t.Format("02 Jan 2006 15:04:05")
}