// Code generated by wit-bindgen-go. DO NOT EDIT.

package store

import (
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// ErrorShape is used for storage in variant or result types.
type ErrorShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(Error{})]byte
}

// OptionListU8Shape is used for storage in variant or result types.
type OptionListU8Shape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(cm.Option[cm.List[uint8]]{})]byte
}

// KeyResponseShape is used for storage in variant or result types.
type KeyResponseShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(KeyResponse{})]byte
}

func lower_OptionU64(v cm.Option[uint64]) (f0 uint32, f1 uint64) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1 := (uint64)(*some)
		f1 = (uint64)(v1)
	}
	return
}