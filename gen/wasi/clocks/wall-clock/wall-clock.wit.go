// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package wallclock represents the imported interface "wasi:clocks/wall-clock@0.2.3".
//
// WASI Wall Clock is a clock API intended to let users query the current
// time. The name "wall" makes an analogy to a "clock on the wall", which
// is not necessarily monotonic as it may be reset.
//
// It is intended to be portable at least between Unix-family platforms and
// Windows.
//
// A wall clock is a clock which measures the date and time according to
// some external reference.
//
// External references may be reset, so this clock is not necessarily
// monotonic, making it unsuitable for measuring elapsed time.
//
// It is intended for reporting the current date and time for humans.
package wallclock

import (
	"go.bytecodealliance.org/cm"
)

// DateTime represents the record "wasi:clocks/wall-clock@0.2.3#datetime".
//
// A time and date in seconds plus nanoseconds.
//
//	record datetime {
//		seconds: u64,
//		nanoseconds: u32,
//	}
type DateTime struct {
	_           cm.HostLayout
	Seconds     uint64
	Nanoseconds uint32
}

// Now represents the imported function "now".
//
// Read the current value of the clock.
//
// This clock is not monotonic, therefore calling this function repeatedly
// will not necessarily produce a sequence of non-decreasing values.
//
// The returned timestamps represent the number of seconds since
// 1970-01-01T00:00:00Z, also known as [POSIX's Seconds Since the Epoch],
// also known as [Unix Time].
//
// The nanoseconds field of the output is always less than 1000000000.
//
//	now: func() -> datetime
//
// [POSIX's Seconds Since the Epoch]: https://pubs.opengroup.org/onlinepubs/9699919799/xrat/V4_xbd_chap04.html#tag_21_04_16
// [Unix Time]: https://en.wikipedia.org/wiki/Unix_time
//
//go:nosplit
func Now() (result DateTime) {
	wasmimport_Now(&result)
	return
}

// Resolution represents the imported function "resolution".
//
// Query the resolution of the clock.
//
// The nanoseconds field of the output is always less than 1000000000.
//
//	resolution: func() -> datetime
//
//go:nosplit
func Resolution() (result DateTime) {
	wasmimport_Resolution(&result)
	return
}