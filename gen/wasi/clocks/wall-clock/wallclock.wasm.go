// Code generated by wit-bindgen-go. DO NOT EDIT.

package wallclock

// This file contains wasmimport and wasmexport declarations for "wasi:clocks@0.2.3".

//go:wasmimport wasi:clocks/wall-clock@0.2.3 now
//go:noescape
func wasmimport_Now(result *DateTime)

//go:wasmimport wasi:clocks/wall-clock@0.2.3 resolution
//go:noescape
func wasmimport_Resolution(result *DateTime)