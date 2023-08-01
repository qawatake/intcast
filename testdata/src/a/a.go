package a

import "a/hoge"

type uint32x uint32

func castInt() {
	var i int
	_ = int(i)         // OK
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // want "unsafe cast"
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // want "unsafe cast"
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"

	_ = hoge.IntToInt8(i) // OK
}

func castInt16() {
	var i int16
	_ = int(i)         // OK
	_ = int16(i)       // OK
	_ = int32(i)       // OK
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // want "unsafe cast"
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // want "unsafe cast"
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castInt32() {
	var i int32
	_ = int(i)         // OK
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // OK
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // want "unsafe cast"
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // want "unsafe cast"
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castInt8() {
	var i int8
	_ = int(i)         // OK
	_ = int16(i)       // OK
	_ = int32(i)       // OK
	_ = int8(i)        // OK
	_ = int64(i)       // OK
	_ = uint(i)        // want "unsafe cast"
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // want "unsafe cast"
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castInt64() {
	var i int64
	_ = int(i)         // OK
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // want "unsafe cast"
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // want "unsafe cast"
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castUint() {
	var i uint
	_ = int(i)         // want "unsafe cast"
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // want "unsafe cast"
	_ = uint(i)        // OK
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castUint16() {
	var i uint16
	_ = int(i)         // OK
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // OK
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // OK
	_ = uint16(i)      // OK
	_ = uint32(i)      // OK
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // OK
	_ = hoge.Uint16(i) // OK
}

func castUint32() {
	var i uint32
	_ = int(i)         // want "unsafe cast"
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // OK
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // OK
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // OK
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castUint64() {
	var i uint64
	_ = int(i)         // want "unsafe cast"
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // want "unsafe cast"
	_ = uint(i)        // OK
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // want "unsafe cast"
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // want "unsafe cast"
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castUint8() {
	var i uint8
	_ = int(i)         // OK
	_ = int16(i)       // OK
	_ = int32(i)       // OK
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // OK
	_ = uint16(i)      // OK
	_ = uint32(i)      // OK
	_ = uint64(i)      // OK
	_ = uint8(i)       // OK
	_ = uint32x(i)     // OK
	_ = hoge.Uint16(i) // OK
}

func castUint32X() {
	var i uint32x
	_ = int(i)         // want "unsafe cast"
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // want "unsafe cast"
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // OK
	_ = uint16(i)      // want "unsafe cast"
	_ = uint32(i)      // OK
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     //  OK
	_ = hoge.Uint16(i) // want "unsafe cast"
}

func castHogeUint16() {
	var i hoge.Uint16
	_ = int(i)         // OK
	_ = int16(i)       // want "unsafe cast"
	_ = int32(i)       // OK
	_ = int8(i)        // want "unsafe cast"
	_ = int64(i)       // OK
	_ = uint(i)        // OK
	_ = uint16(i)      // OK
	_ = uint32(i)      // OK
	_ = uint64(i)      // OK
	_ = uint8(i)       // want "unsafe cast"
	_ = uint32x(i)     // OK
	_ = hoge.Uint16(i) // OK

	type uint16x uint16
	_ = uint16x(i) // OK
}

func lintIgnore() {
	var i uint64
	//lint:ignore intcast reason
	_ = int(i) // OK
	//lint:ignore intcast
	_ = int(i) // want "unsafe cast"
	//lint:ignore hoge reason
	_ = int(i) // want "unsafe cast"
	//lint:ignore hoge
	_ = int(i) // want "unsafe cast"
}
