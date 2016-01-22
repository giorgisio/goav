// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/opt.h>
import "C"
import "unsafe"

func AvOptSetBin(obj unsafe.Pointer, name string, val * uint8, size int, search_flags int) int {
	return int(C.av_opt_set_bin(obj, C.CString(name), (*C.uint8_t)(val), C.int(size), C.int(search_flags)))
}