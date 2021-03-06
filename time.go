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

// SQLStr formates to MySQL time string `02 Jan 2006 15:04:05`
func (t Time)SQLStr() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// BeginningOfMonth return the begin of the month of t
func (t Time)BeginningOfMonth() time.Time {
	year, month, _ := time.Time(t).Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// TimeEndOfMonth return the end of the month of t
func (t Time)EndOfMonth() time.Time {
	return t.BeginningOfMonth().AddDate(0, 1, -1)
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