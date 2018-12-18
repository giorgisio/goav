// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	AV_PIX_FMT_YUV   = int(C.AV_PIX_FMT_YUV)
	AV_PIX_FMT_RGB24 = int(C.AV_PIX_FMT_RGB24)
	AV_PIX_FMT_RGBA  = int(C.AV_PIX_FMT_RGBA)
)

func (pf PixelFormat) String() string {
	switch int(pf) {
	case AV_PIX_FMT_YUV; return "YUV"
	case AV_PIX_FMT_RGB24: return "RGB24"
	case AV_PIX_FMT_RGBA: return "RGBA"
	}
}